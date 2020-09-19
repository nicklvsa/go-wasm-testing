if (!('WebAssembly' in window)) {
    alert('You must use a browser that supports WebAssebmly!');
}

const go = new Go();

const loadWebAssembly = async (filename, imports) => {
    return fetch(filename).then((response) => response.arrayBuffer()).then((buffer) => {

        imports = imports || {};
        imports.env = imports.env || {}; 
        imports.env.memoryBase = imports.env.memoryBase || 0;
        imports.env.tableBase = imports.env.tableBase || 0;

        if (!imports.env.memory) {
            imports.env.memory = new WebAssembly.Memory({initial: 256});
        }

        if (!imports.env.table) {
            imports.env.table = new WebAssembly.Table({initial: 0, element: 'anyfunc'});
        }

        return WebAssembly.instantiate(buffer, Object.assign(imports, {...go.importObject}));
    });
};

loadWebAssembly('../main.wasm').then((wasm) => {
    const exampleBtn = document.querySelector('#example_btn');
    exampleBtn.disabled = false;
    exampleBtn.addEventListener('click', () => {
        go.run(wasm.instance);
    });
});