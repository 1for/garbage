;;(add-to-list 'load-path "~/.emacs.d/")

;;加载基本配置

;; Added by Package.el.  This must come before configurations of
;; installed packages.  Don't delete this line.  If you don't want it,
;; just comment it out by adding a semicolon to the start of the line.
;; You may delete these explanatory comments.
(load "~/.emacs.d/config/base-config.el")
;;加载linum插件配置
(load "~/.emacs.d/config/linum-config.el")
;;加载tabbar插件配置
;;(load "~/.emacs.d/config/tabbar-config.el")
;;加载js3插件配置
;;(load "~/.emacs.d/config/js3-config.el")
;;(load "~/.emacs.d/config/js2-config.el")
;;加载markdown-mode插件配置
;;(load "~/.emacs.d/config/markdown-mode-config.el")
;;加载google-translate插件配置
;;(load "~/.emacs.d/config/google-translate-config.el")
;;(load "~/.emacs.d/config/baidu-translate-config.el")
;;加载ace-jump-line-mode
(load "~/.emacs.d/config/ace-jump-line-mode-config.el")
;;加载switch-window插件配置
(load "~/.emacs.d/config/switch-window-config.el")
;;加载yasnippet插件配置
;;(load "~/.emacs.d/config/yasnippet-config.el")
;;加载auto-complete插件配置
;;(load "~/.emacs.d/config/ac-config.el")
;;加载web-mode插件配置
;;(load "~/.emacs.d/config/web-mode-config.el")
;;加载minimap插件配置
;;(load "~/.emacs.d/config/minimap-config.el")
;;加载idle-highlight-mode
;;(load "~/.emacs.d/config/idle-highlight-mode-config.el")
;;加载dirtree插件配置
(load "~/.emacs.d/config/dirtree-config.el")
;;加载Nyan mode
;;(load "~/.emacs.d/config/nyan-mode-config.el")
;;加载p4
;;(load "~/.emacs.d/config/p4-config.el")
;;加载smex
(load "~/.emacs.d/config/smex-config.el")
(load "~/.emacs.d/config/recent-jump.el")
(load "~/.emacs.d/config/projectile-conf.el")

;; 使用 gopls 替代
;;(load "~/.emacs.d/config/gotags-config.el")
;;(load "~/.emacs.d/config/gocode-config.el") ;; https://github.com/nsf/gocode
(load "~/.emacs.d/config/godev-config.el")

;;加载php-mode插件配置
(load "~/.emacs.d/config/php-mode-config.el")

(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(package-selected-packages
   '(company lsp-ui lsp-mode ac-js2 js2-mode magit flycheck solidity-mode yaml-mode request helm-projectile project popup phpunit multi-web-mode go-mode find-file-in-repository find-file-in-project fill-column-indicator dsvn autothemer all-the-icons))
 '(web-mode-code-indent-offset 4)
 '(web-mode-css-indent-offset 4)
 '(web-mode-markup-indent-offset 4)
 '(web-mode-script-padding 4)
 '(web-mode-style-padding 4))

(custom-set-faces
 ;; custom-set-faces was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 )
