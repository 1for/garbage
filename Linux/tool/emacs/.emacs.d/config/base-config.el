;;no backup files
(setq make-backup-files nil)

;;no start info
(setq inhibit-startup-message t)

;; theme
(add-to-list 'custom-theme-load-path "~/.emacs.d/themes/moe")


;;默认选择moe主题
(load-theme 'moe-dark t)

;;emacs包管理
(require 'package)
;;(add-to-list 'package-archives 
;;			 '("marmalade" . "http://marmalade-repo.org/packages/") t)
(add-to-list 'package-archives
			 '("melpa" . "https://melpa.org/packages/") t)
;;(add-to-list 'package-archives '("melpa" . "http://melpa.milkbox.net/packages/"))
(package-initialize)

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

;;字体设置
;;(set-default-font "-outline-Lucida Console-normal-normal-normal-mono-*-*-*-*-*-*-iso10646-1" 1)

;;设置光标
(setq-default cursor-type 'bar)

;;文件编码设置
(prefer-coding-system 'utf-8)
(setq-default indent-tabs-mode  nil)

;;(setq file-name-coding-system 'gbk)  ;;仅window下启用
(setq default-buffer-file-coding-system 'utf-8)

;;自定义快捷键
;;(global-set-key (kbd "<f5>")   'find-file-in-repository)
;;(global-set-key (kbd "<f6>")   'phpunit-current-class)
(global-set-key (kbd "<f7>")   'helm-projectile-find-file)
(global-set-key (kbd "<f8>")   'helm-projectile-grep)



;;hack
;; auto insert space
;;(global-set-key (kbd ",")
;;		#'(lambda ()
;;		    (interactive)
;;		    (insert ", ")))
