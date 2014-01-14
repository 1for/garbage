;;web-mode插件配置
(add-to-list 'load-path "~/.emacs.d/plugin/web-mode")

(require 'web-mode)
(add-to-list 'auto-mode-alist '("\\.html?\\'" . web-mode))
