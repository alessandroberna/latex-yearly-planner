# latex-yearly-planner-amoled
PDF planner designed for e-ink devices, edited to be used with amoled displays.

See [discussions](https://github.com/kudrykv/latex-yearly-planner/discussions) for available planners and their variations.

Pre-generated planners are available [here (amoled)](https://github.com/kudrykv/latex-yearly-planner/discussions/134) or [here (original)](https://github.com/kudrykv/latex-yearly-planner/discussions/57).

## Planner Generation:
### Install Dependencies
1. [Go Language](https://go.dev/dl/)
2. [LaTex](https://miktex.org/download) & [PDFLaTeX](https://www.latex-project.org/get/)
3. From the project directory, run the following command after updating
 'PLANNER_YEAR' below. This should generate the PDF in the 'out' directory.
<code>PLANNER_YEAR=2022 \
PASSES=1 \
CFG="cfg/base.yaml,cfg/template_breadcrumb.yaml,cfg/sn_a5x.breadcrumb.default.yaml" \
NAME="sn_a5x.breadcrumb.default" \
./single.sh</code> 
 
   [Source](https://github.com/kudrykv/latex-yearly-planner/discussions/34#discussioncomment-3128344)

4. Check the "out" directory for the 'pdf' planner. To move it to your device
, follow the manufacturer's instructions on how to load a PDF on your device.

If you encounter any problems related to '.sty' files you likely need to
 install some Latex related dependencies. Copying the error and search using
  your favorite search engine should get you on track to resolving the
   dependency issues. All the best!

### Alternative install
Instead of installing the dependencies manually, this repository is defined as a Nix flake which specifies fixed versions of all the required dependencies. 

1. [Install Nix](https://nixos.org/download.html)
2. Build a planner pdf using `nix build`
3. Or, if you want to develop the code, enter a shell with all the dependencies present using `nix develop`
