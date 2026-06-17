#!/bin/bash
set -eu -o pipefail
# Parameters
imageLocation=""
imageUrl=""
targetFolder="macadam-test"
macadamBinary=""
ginkgoLabelFilter=""
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -imageLocation)
        imageLocation="$2"
        shift 
        shift 
        ;;
        -imageUrl)
        imageUrl="$2"
        shift 
        shift 
        ;;
        -targetFolder)
        targetFolder="$2"
        shift 
        shift 
        ;;
        -macadamBinary)
        macadamBinary="$2"
        shift 
        shift 
        ;;
        -ginkgoLabelFilter)
        ginkgoLabelFilter="$2"
        shift 
        shift 
        ;;
        *)    # unknown option
        shift 
        ;;
    esac
done

# Prepare results folder
mkdir -p $HOME/$targetFolder/results

if [ -n "$macadamBinary" ]; then
  cp "$macadamBinary" "$HOME/$targetFolder/bin/macadam"
fi
macadamBinary="$HOME/$targetFolder/bin/macadam"

if [ -z "$imageLocation" ] && [ -n "$imageUrl" ]; then
    filename="$(basename "$imageUrl")"
    echo "Download image $filename ..."
    downloadImage="$HOME/$targetFolder/testdata/$filename"
    mkdir -p "$HOME/$targetFolder/testdata"

    curl -L "$imageUrl" -o "$downloadImage"
    lower_filename="$(echo "$filename" | tr '[:upper:]' '[:lower:]')"
    
    if [[ "$lower_filename" == *.raw || \
         "$lower_filename" == *.qcow2 ]]; then
        imageLocation="$downloadImage"
    else
        echo "not supported format $filename"
        exit 1
    fi
fi

OS="$(uname)"

if [ "$OS" = "Linux" ]; then
    oslabel="linux"
elif [ "$OS" = "Darwin" ]; then
    oslabel="darwin"
else
    echo "unsupported OS $OS"
    exit 1
fi

if [ -z "$ginkgoLabelFilter" ]; then
  ginkgoLabelFilter="$oslabel"
else
  ginkgoLabelFilter="$oslabel&&$ginkgoLabelFilter"
fi

cd $HOME/$targetFolder/bin
set +e
./e2e.test \
  --macadam-binary="$macadamBinary" \
  --ginkgo.label-filter="$ginkgoLabelFilter" \
  --image="$imageLocation" \
  2>&1 | tee e2e.result
test_exit=${PIPESTATUS[0]}

cd ..
cp -f bin/out/* results/ 2>/dev/null
cp -f bin/e2e.result results/ 2>/dev/null

exit "$test_exit"