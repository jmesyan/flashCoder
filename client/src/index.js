 function component() {
    var element = document.createElement('div');

    // Lodash, now imported by this script
    element.innerHTML = "hello react client";

    return element;
}

 document.body.appendChild(component());