     
     if (! WebAssembly.instantiateStreaming){
	 
         WebAssembly.instantiateStreaming = async (resp, importObject) => {
             const source = await (await resp).arrayBuffer();
             return await WebAssembly.instantiate(source, importObject);
         };
     }

     const go = new Go();

     let mod, inst;

     WebAssembly.instantiateStreaming(fetch("chicken.wasm"), go.importObject).then(
	 
          async result => {
	      document.getElementById("button").innerText = "Convert";
	      document.getElementById("button").removeAttribute("disabled");
              mod = result.module;
              inst = result.instance;
	      await go.run(inst);
          }
      );

     async function convert() {

	 var raw_el = document.getElementById("raw");
	 // var chicken_el = document.getElementById("chicken");

	 var raw_txt = raw_el.value;
	 var chicken_txt = chicken("zxx", raw_txt);

	 raw_el.value = chicken_txt;
     }