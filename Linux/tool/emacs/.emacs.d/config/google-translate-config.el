(add-to-list 'load-path "~/.emacs.d/plugin/google-translate")

(require 'google-translate)
(setq google-translate-default-target-language "zh-CN")
(setq google-translate-default-source-language "en")
(global-set-key "\C-ct" 'google-translate-at-point)
(global-set-key "\C-cT" 'google-translate-query-translate)

