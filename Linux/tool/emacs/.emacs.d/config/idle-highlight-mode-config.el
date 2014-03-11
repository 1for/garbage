;;idle-highlight-mode插件配置
(add-to-list 'load-path "~/.emacs.d/plugin/idle-highlight-mode")

(autoload 'idle-highlight-mode "idle-highlight-mode" "highlight the word the point is on" t)
(add-hook 'find-file-hook 'idle-highlight-mode)
