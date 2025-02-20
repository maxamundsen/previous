// Allows you to select elements relative to self.
// Ex:
// <div id="test-element">
//     <button>Hello!</button>
// </div>
//
// <script>
//     var testEl = document.getElementById("test-element");
//     var theButton = testEl.querySelectorRelative("button"); // this selects the button contained inside `testEl`
// </script>
Element.prototype.querySelectorRelative = function(selector){
    // Adding a custom attribute to refer for selector
    this.setAttribute('data-unique-id', '1');

    // Replace "this " string with custom attribute's value
    // You can also add a unique class name instead of adding custom attribute
    selector = '[data-unique-id="1"] ' + selector

    // Get the relative element
    var relativeElement = document.querySelector(selector);

    // After getting the relative element, the added custom attribute is useless
    // So, remove it
    this.removeAttribute('data-unique-id');

    // return the fetched element
    return relativeElement;
}