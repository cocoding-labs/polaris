forge build --extra-output-files bin --extra-output-files abi --root .

mkdir -p ./bindings/forma/precompile/nativeminter
abigen --pkg nativeminter \
  --abi ./out/INativeMinter.sol/INativeMinter.abi.json \
  --bin ./out/INativeMinter.sol/INativeMinter.bin \
  --out ./bindings/forma/precompile/nativeminter/i_nativeminter.abigen.go \
  --type NativeMinter

mkdir -p ./bindings/forma/precompile/compress
abigen --pkg compress \
  --abi ./out/ICompress.sol/ICompress.abi.json \
  --bin ./out/ICompress.sol/ICompress.bin \
  --out ./bindings/forma/precompile/compress/i_compress.abigen.go \
  --type Compress

mkdir -p ./bindings/forma/precompile/jsonutil
abigen --pkg jsonutil \
  --abi ./out/IJsonUtil.sol/IJsonUtil.abi.json \
  --bin ./out/IJsonUtil.sol/IJsonUtil.bin \
  --out ./bindings/forma/precompile/jsonutil/i_jsonutil.abigen.go \
  --type JsonUtil
