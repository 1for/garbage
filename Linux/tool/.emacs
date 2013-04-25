(setq make-backup-files nil)

;; Config cedet Begin ------
(add-to-list 'load-path "/home/liuyang/cedet-1.0/common")
(require 'cedet)
(require 'semantic-ia)

;; Enable EDE (Project Management) features
(global-ede-mode 1)
  
(semantic-load-enable-excessive-code-helpers)
(semantic-load-enable-semantic-debugging-helpers)
  
;; Enable SRecode (Template management) minor-mode.
(global-srecode-minor-mode 1) 

(global-ede-mode t)
;; Config cedet End ------

;; Config ecb Begin ------
(add-to-list 'load-path "~/.emacs.d/ecb")

(require 'ecb)
(setq ecb-auto-activate t)
(custom-set-variables
  ;; custom-set-variables was added by Custom.
  ;; If you edit it by hand, you could mess it up, so be careful.
  ;; Your init file should contain only one such instance.
  ;; If there is more than one, they won't work right.
 '(ecb-options-version "2.40"))
(custom-set-faces
  ;; custom-set-faces was added by Custom.
  ;; If you edit it by hand, you could mess it up, so be careful.
  ;; Your init file should contain only one such instance.
  ;; If there is more than one, they won't work right.
 )
