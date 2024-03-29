(add-to-list 'load-path "~/.emacs.d/plugin/php-mode")

(require 'use-package)
(require 'php-mode)
(setq default-tab-width 4)
(setq indent-tabs-mode nil)
(setq c-default-style "bsd") ; set code style bsd
(setq-default c-basic-offset 4)

;;限制每行字数
;;(add-hook 'php-mode-hook 'fci-mode)

(add-hook 'php-mode-user-hook 'turn-on-font-lock)
(add-hook 'php-mode-hook
          '(lambda () (define-abbrev php-mode-abbrev-table "ex" "extends")))

;;自动补全设置
;;(setq php-manual-path "/home/liuyang/doc/php-manual/")
;; (global-set-key [(control tab)] 'php-complete-function)

;;(use-package lsp-mode
;; :config
;; (setq lsp-prefer-flymake nil)
;; :hook (php-mode . lsp)
;; :commands lsp)
