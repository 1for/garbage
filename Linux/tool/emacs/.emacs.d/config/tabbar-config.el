;;tabbar插件配置
(add-to-list 'load-path "~/.emacs.d/plugin/tabbar")

(require 'tabbar)
(tabbar-mode t)
(global-set-key [27 up] 'tabbar-forward-group)
(global-set-key [27 down] 'tabbar-backward-group)
(global-set-key [27 left] 'tabbar-backward)
(global-set-key [27 right] 'tabbar-forward)

