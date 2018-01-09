(require 'url)
(require 'json)

(defvar baidu-appid "20171218000106053")
(defvar baidu-seckey "MRtxS9OhhWEejuh_46LM")
(defvar baidu-translate-baseurl "http://api.fanyi.baidu.com/api/trans/vip/translate")

(defun baidu-translate-at-point(prefix)
  "baidu translate"
  (interactive "P")
  (let* ((query-url
          (baidu-build-url
           (if (use-region-p)
               (buffer-substring-no-properties (region-beginning) (region-end))
             (or (current-word t)
                 (error "No word at point."))))))
    (message (cdr (car (aref (cdr (car (baidu-response-json query-url))) 0))))))

(defun baidu-response-json(url)
  (json-read-from-string
   (with-current-buffer (url-retrieve-synchronously url)
     (set-buffer-multibyte t)
     (goto-char (point-min))
     (re-search-forward (format "\n\n"))
     (delete-region (point-min) (point))
     (prog1 (buffer-string)))))
    
(defun baidu-build-url(query)
  "build url for request"
  (let* ((salt (number-to-string (random 1000000)))
         (from "en")
         (to "zh")
         (sign (baidu-build-sign query baidu-appid salt baidu-seckey)))
    (concat baidu-translate-baseurl "?q=" query "&appid=" baidu-appid
            "&salt=" salt "&from=" from "&to=" to "&sign=" sign)))

(defun baidu-build-sign(query appid salt seckey)
  (md5 (concat appid query salt seckey)))

(provide 'baidu-translate)
