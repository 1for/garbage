;;smex²å¼şÅäÖÃ
(add-to-list 'load-path "~/.emacs.d/plugin/smex")

(require 'smex)

(global-set-key (kbd "M-x") 'smex)
(global-set-key (kbd "M-X") 'smex-major-mode-commands)
