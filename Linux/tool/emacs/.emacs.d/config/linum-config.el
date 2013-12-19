;;linum+插件配置
(add-to-list 'load-path "~/.emacs.d/plugin/linum+")

;; 显示行号,列号
(global-linum-mode t)                                                                                                       
(setq linum-format "%d |")
(setq column-linum-mode t)

(require 'linum+)
