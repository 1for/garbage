;; emacs 24.5 raise error
;;(add-to-list 'load-path "~/.emacs.d/")

;;加载基本配置
(load "~/.emacs.d/config/base-config.el")
;;加载linum插件配置
(load "~/.emacs.d/config/linum-config.el")
;;加载tabbar插件配置
(load "~/.emacs.d/config/tabbar-config.el")
;;加载js3插件配置
(load "~/.emacs.d/config/js3-config.el")
;;加载php-mode插件配置
(load "~/.emacs.d/config/php-mode-config.el")
;;加载markdown-mode插件配置
(load "~/.emacs.d/config/markdown-mode-config.el")
;;加载google-translate插件配置
(load "~/.emacs.d/config/google-translate-config.el")
;;加载ace-jump-line-mode
(load "~/.emacs.d/config/ace-jump-line-mode-config.el")
;;加载switch-window插件配置
(load "~/.emacs.d/config/switch-window-config.el")
;;加载yasnippet插件配置
(load "~/.emacs.d/config/yasnippet-config.el")
;;加载auto-complete插件配置
(load "~/.emacs.d/config/ac-config.el")
;;加载web-mode插件配置
(load "~/.emacs.d/config/web-mode-config.el")
;;加载minimap插件配置
(load "~/.emacs.d/config/minimap-config.el")
;;加载idle-highlight-mode
(load "~/.emacs.d/config/idle-highlight-mode-config.el")
;;加载dirtree插件配置
(load "~/.emacs.d/config/dirtree-config.el")
;;加载Nyan mode
(load "~/.emacs.d/config/nyan-mode-config.el")
;;加载p4
(load "~/.emacs.d/config/p4-config.el")
;;加载smex
(load "~/.emacs.d/config/smex-config.el")

(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(custom-safe-themes (quote ("9fd20670758db15cc4d0b4442a74543888d2e445646b25f2755c65dcd6f1504b" default)))
 '(js3-curly-indent-offset 2)
 '(js3-expr-indent-offset 2)
 '(js3-lazy-commas t)
 '(js3-lazy-dots t)
 '(js3-lazy-operators t)
 '(js3-paren-indent-offset 2)
 '(js3-square-indent-offset 2)
 '(menu-bar-mode nil)
 '(truncate-partial-width-windows nil))
(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(default ((t (:family "YaHei Consolas Hybrid" :foundry "outline" :slant normal :weight normal :height 105 :width normal)))))
