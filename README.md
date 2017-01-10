## Mergepdf

Walks through file system tree and merges pdf documents in 'leaf' directories

```shell
  $ go run main.go <root directory> <path to pdfbox.jar> <output file name>
```

### Jar file info:
This repo contains the pdfbox.jar file already, but if you want to download it yourself get it from apache [here](https://pdfbox.apache.org/download.cgi)

### Example folder structure:
 - root_folder
  - sub_folder
    - pdf_a.pdf
    - pdf_b.pdf
  - sub_folder
    - pdf_c.pdf
    - pdf_d.pdf
  - sub_folder
    - pdf_e.pdf
    - pdf_f.pdf

### Example command:
```shell
  $ go run main.go ./root_folder ./jar/pdfbox.jar output.pdf
```

### Example output folder structure:
 - root_folder
  - sub_folder
    - pdf_a.pdf
    - pdf_b.pdf
    - output.pdf
  - sub_folder
    - pdf_c.pdf
    - pdf_d.pdf
    - output.pdf
  - sub_folder
    - pdf_e.pdf
    - pdf_f.pdf
    - output.pdf
