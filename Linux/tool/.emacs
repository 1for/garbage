(add-to-list 'load-path "~/.emacs.d/")

;;no backup files
(setq make-backup-files nil)

;;no start info
(setq inhibit-startup-message t)

;;set color theme
(set-foreground-color "green")
(set-background-color "black")
(set-cursor-color "gold")
(set-mouse-color "gold")

;; BEGIN js3 --
;;;Add the following custom-set-variables to use 'lazy' modes
(custom-set-variables
  ;; custom-set-variables was added by Custom.
  ;; If you edit it by hand, you could mess it up, so be careful.
  ;; Your init file should contain only one such instance.
  ;; If there is more than one, they won't work right.
 '(js3-lazy-commas t)
 '(js3-lazy-operators t)
 '(js3-lazy-dots t)
 '(js3-expr-indent-offset 2)
 '(js3-paren-indent-offset 2)
 '(js3-square-indent-offset 2)
 '(js3-curly-indent-offset 2)
)

(autoload 'js3-mode "js3" nil t)
(add-to-list 'auto-mode-alist '("\\.js$" . js3-mode))
;; END js3 --

;; BEGIN tabbar
(require 'tabbar)
(tabbar-mode t)
(global-set-key [27 up] 'tabbar-forward-group)
(global-set-key [27 down] 'tabbar-backward-group)
(global-set-key [27 left] 'tabbar-backward)
(global-set-key [27 right] 'tabbar-forward)
;; END tabbar

;; BEGIN php-mode
(require 'php-mode)
;; END php-mode

;; do not use tab,instead of 4 space
(setq tab-width 4)
(setq c-basic-offset 4)
(custom-set-variables
  ;; custom-set-variables was added by Custom.
  ;; If you edit it by hand, you could mess it up, so be careful.
  ;; Your init file should contain only one such instance.
  ;; If there is more than one, they won't work right.
 '(indent-tabs-mode nil))
(custom-set-faces
  ;; custom-set-faces was added by Custom.
  ;; If you edit it by hand, you could mess it up, so be careful.
  ;; Your init file should contain only one such instance.
  ;; If there is more than one, they won't work right.
 )
