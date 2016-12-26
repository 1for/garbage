(add-to-list 'load-path "~/.emacs.d/plugin/recent-jump")

(require 'recent-jump)
(global-set-key (kbd "C-c h") 'recent-jump-backward)
(global-set-key (kbd "C-c j") 'recent-jump-forward)

(define-globalized-minor-mode global-recent-jump-mode recent-jump-mode
    (lambda () (recent-jump-mode 1)))

(global-recent-jump-mode 1)
