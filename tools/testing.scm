;;; -*- mode: Scheme -*-

;;; Copyright 2015 SteelSeries ApS. All rights reserved.
;;; Use of this source code is governed by a BSD-style
;;; license that can be found in the LICENSE file.

(define number-of-passes 0)
(define number-of-failures 0)
(define number-of-errors 0)

(define failure-messages '())
(define error-messages '())

(define verbose-tests nil)
(define context-name "")
(define it-name "")


(define (reset-testing )
  (set! number-of-passes 0)
  (set! number-of-failures 0)
  (set! number-of-errors 0)
  (set! failure-messages '())
  (set! error-messages '())
  (set! verbose-tests nil))

(define (log-pass msg)
  (set! number-of-passes (succ number-of-passes))
  (when verbose-tests
    (format #t "    ~A~%" msg)))


(define (log-failure prefix msg)
  (set! number-of-failures (succ number-of-failures))
  (let ((failure-message (format #f "~A ~A:~%    ~A~%      - ~A" context-name it-name prefix msg)))
    (set! failure-messages (cons failure-message failure-messages))
    (when verbose-tests
          (format #t "    ~A~%      - ~A~%" prefix msg))))

(define (log-error err)
  (set! number-of-errors (succ number-of-errors))
  (let ((error-message (format #f "~A ~A:~%    ERROR: ~A" context-name it-name err)))
    (set! error-messages (cons error-message error-messages))
    (when verbose-tests
          (format #t "    ERROR: ~A~%" err))))

(defmacro (context label setup . body)
  (if (not (or (symbol? label) (string? label)))
      (error "The label of a describe must be a symbol or string.")
      `(begin (when verbose-tests
                (format #t "~%~A~%" ,label))
              (set! context-name ,label)
              (for-each (lambda (it-clause)
                          ,@setup
                          (eval it-clause))
                        ',body))))

(defmacro (it label . body)
  (if (not (or (symbol? label) (string? label)))
      (error "The label of a describe must be a symbol or string.")
      `(begin (when verbose-tests
                (format #t "~%  ~A~%" ,label))
              (set! it-name ,label)
              (on-error (begin ,@body)
                        (lambda (err)
                          (let* ((err-parts (string-split err "\n"))
                                 (last-line (car (last-pair err-parts)))
                                 (report last-line))
                            (log-error report)))))))

(defmacro (assert-true sexpr)
  `(let ((actual ,sexpr)
         (msg (format #f "(assert-true ~S)" ',sexpr)))
     (if actual
         (log-pass msg)
         (log-failure msg "expected true, but was false"))))


(defmacro (assert-false sexpr)
  `(let ((actual ,sexpr)
         (msg (format #f "(assert-false ~S)" ',sexpr)))
     (if (not actual)
         (log-pass msg)
         (log-failure msg "expected false, but was true"))))


(defmacro (assert-nil sexpr)
  `(let ((actual ,sexpr)
         (msg (format #f "(assert-null ~S)" ',sexpr)))
     (if (nil? actual)
         (log-pass msg)
         (log-failure msg "expected nil, but wasn't"))))


(defmacro (assert-not-nil sexpr)
  `(let ((actual ,sexpr)
         (msg (format #f "(assert-not-null ~S)" ',sexpr)))
     (if (not (nil? actual))
         (log-pass msg)
         (log-failure msg "expected not nil, but was"))))


(defmacro (assert-eq sexpr expected-sexpr)
  `(let* ((actual ,sexpr)
          (expected ,expected-sexpr)
          (msg (format #f "(assert-eq ~S ~S)" ',sexpr ',expected-sexpr)))
     (if (equal? actual expected)
         (log-pass msg)
         (log-failure msg (format #f "expected ~S, but was ~S" expected actual)))))


(defmacro (assert-neq sexpr expected-sexpr)
  `(let* ((actual ,sexpr)
          (expected ,expected-sexpr)
          (msg (format #f "(assert-neq ~S ~S)" ',sexpr ',expected-sexpr)))
     (if (not (equal? actual expected))
         (log-pass msg)
         (log-failure msg (format #f "did not expect ~S, but it was" expected)))))

(defmacro (assert-error **sexpr**)
  `(let ((msg (format #f "(assert-error ~S)" ',**sexpr**)))
     (on-error ,**sexpr**
               (lambda (err)
                 (log-pass msg))
               (lambda ()
                 (log-failure msg "expected an error, but there wasn't")))))

(defmacro (assert-nerror **sexpr**)
  `(let ((msg (format #f "(assert-nerror ~S)" ',**sexpr**)))
     (on-error ,**sexpr**
               (lambda (err)
                 (log-failure msg (format #f "expected no error, but error was ~A" err)))
               (lambda ()
                 (log-pass msg)))))

(defmacro (assert-memq sexpr object)
  `(let* ((searched-for ,object)
          (result (memq ,object ,sexpr))
          (msg (format #f "(assert-memq ~S ~S)" ',sexpr ',object)))
     (if result
         (log-pass msg)
         (log-failure msg (format #f "expected ~S to contain ~S, but it didn't" ',sexpr searched-for)))))

(define (dump-summary duration)
  (format #t "~%Ran ~A tests in ~A seconds~%"
          (+ number-of-passes number-of-failures number-of-errors)
          (/ duration 1000.0))
  (format #t "~A pass~A, ~A failure~A, ~A error~A~%"
          number-of-passes (if (eq? number-of-passes 1) "" "es")
          number-of-failures (if (eq? number-of-failures 1) "" "s")
          number-of-errors (if (eq? number-of-errors 1) "" "s"))
  (unless (zero? number-of-failures)
          (format #t "~%Failure~A:~%" (if (eq? number-of-failures 1) "" "s"))
    (for-each (lambda (m) (format #t "  ~A~%" m))
              failure-messages))
  (unless (zero? number-of-errors)
    (format #t "~%Error~A:~%" (if (eq? number-of-errors 1) "" "s"))
    (for-each (lambda (m) (format #t "  ~A~%" m))
              error-messages))
  (and (zero? number-of-failures) (zero? number-of-errors)))

(define (run-all-tests test-dir . optionals)
  (reset-testing)
  (set! verbose-tests (not (nil? optionals)))
  (let ((t (time (for-each (lambda (filename)
                             (when verbose-tests (format #t "~%------~%Loading: ~A~%" filename))
                             (load filename))
                           (list-directory test-dir "*_test.scm")))))
    (dump-summary t)))

(define (run-test test-file . optionals)
  (reset-testing)
  (set! verbose-tests (not (nil? optionals)))
  (let ((t (time (load test-file))))
    (dump-summary t)))