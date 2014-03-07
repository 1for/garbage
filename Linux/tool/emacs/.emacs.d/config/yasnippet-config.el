;;yasnippet插件配置
(add-to-list 'load-path "~/.emacs.d/plugin/yasnippet")

(require 'yasnippet)
(yas/initialize)
(yas/load-directory "~/.emacs.d/plugin/yasnippet/snippets")

