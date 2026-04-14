(defun dev-run8 ()
  "Run a script when a PHP file is saved."
  (interactive)
  (when (and (bound-and-true-p buffer-file-name)
             (string-suffix-p ".php" buffer-file-name))
    (shell-command (format "dev-run8 %s" buffer-file-name))))

;;(add-hook 'after-save-hook 'dev-run8)
