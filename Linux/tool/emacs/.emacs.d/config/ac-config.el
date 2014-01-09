;;auto-complete插件配置
(add-to-list 'load-path "~/.emacs.d/plugin/ac")

(require 'auto-complete-config)
(ac-config-default)
(add-to-list 'ac-dictionary-directories "~/.emacs.d/plugin/ac/dict")
