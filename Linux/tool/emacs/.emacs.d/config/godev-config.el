;; Enable company-mode for auto-completion
(add-hook 'after-init-hook 'global-company-mode)

;; Enable LSP mode
(require 'lsp-mode)
(add-hook 'go-mode-hook #'lsp-deferred)

;; Configure LSP UI
(require 'lsp-ui)
(add-hook 'lsp-mode-hook 'lsp-ui-mode)
