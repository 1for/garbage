;;no backup files
(setq make-backup-files nil)

;;no start info
(setq inhibit-startup-message t)

;;set color theme
(set-foreground-color "green")
(set-background-color "black")
(set-cursor-color "gold")
(set-mouse-color "gold")
(add-to-list 'custom-theme-load-path "~/.emacs.d/themes/molokai")

;;emacs包管理
(require 'package)
(add-to-list 'package-archives 
	'("marmalade" . "http://marmalade-repo.org/packages/") t)

;;隐藏菜单栏
(menu-bar-mode -1)

;;隐藏^M
(defun remove-dos-eol ()
  "Do not show ^M in files containing mixed UNIX and DOS line endings."
  (interactive)
  (setq buffer-display-table (make-display-table))
  (aset buffer-display-table ?\^M []))

;;打开ido
(ido-mode t)
