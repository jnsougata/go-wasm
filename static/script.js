const goWasm = new Go();

WebAssembly.instantiateStreaming(fetch("../static/main.wasm"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance)

        

    })

