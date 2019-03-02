# latex-scribe

The code in this repo powers a Firefox extension that parses handwriting into LaTeX symbols.
The extension can insert the recognized text into the editor at [Overleaf.com](https://www.overleaf.com) directly.

## Installation

- [Click here to install the browser extension](https://github.com/dwetterau/latex-scribe/raw/master/published/latex_scribe-0.2-fx.xpi). 
- When the download modal pops up, select "Open With" Firefox
- Firefox will then ask if you want to add the extension

You can find  other versions to install in [this directory of the repo](https://github.com/dwetterau/latex-scribe/tree/master/published).

That's it.

## How to use it
1. Follow the steps in the installation sectionr.
2. Visit [Overleaf.com](https://www.overleaf.com) and begin editing a file.
3. A little white box will show up in the bottom of the editor.
4. Type whatever you'd like, then place your cursor where you want to insert handwritten LaTeX.
5. Draw your latex in the white box, then press "Submit".
6. Clear out your drawing by pressing "Clear".

## How does it work?

- There's a server running at [https://latex.davidw.tech](https://latex.davidw.tech) that receives the input drawings from the browser extension. 
- These drawings are then processed with the [mathpix API](https://docs.mathpix.com/) and the output is returned to the extension.
- The extension then inserts the output into the Overleaf editor directly (via the ACE editor that they're using).

## Questions, Comments, Concerns?

Reach out to me at the email address mentioned on [my website](https://davidw.tech).
