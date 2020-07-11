const go = new Go();
WebAssembly.instantiateStreaming(fetch("tests.wasm"), go.importObject).then((result) => {
  go.run(result.instance);
});
