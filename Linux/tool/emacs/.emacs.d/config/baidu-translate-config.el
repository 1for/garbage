(add-to-list 'load-path "~/.emacs.d/plugin/baidu-translate")

(require 'baidu-translate)
(global-set-key "\C-ct" 'baidu-translate-at-point)


