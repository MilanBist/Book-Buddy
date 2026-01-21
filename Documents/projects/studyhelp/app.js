const button = document.querySelector('#sidebar_logo');
let isopen = false;
let new_div = null;
button.addEventListener('click', ()=>{

    // get each of the icons from the div and print its value

    // new_div.style.height = '2em';
    // new_div.style.width = '80%'
    // new_div.style.backgroundColor = 'yellow';
    // new_div.style.marginLeft = 'auto';
    // new_div.style.marginRight = 'auto';
    // new_div.style.display = 'flex';
    // new_div.style.flexDirection = 'row';
    if(!isopen){
        const body = document.querySelector('#main_body #body h2');
        const icons = document.querySelectorAll("#sidebar i");
        // get the mouse click option and set the logo's background to green
        const logo = document.querySelector('#main_body #sidebar_logo');
        logo.style.backgroundColor = 'green';
        console.log(icons)
        // I have got the icons then view all the icons in the screen
        // make a div at the top of the of the main body
        new_div = document.createElement('div');
        body.append(new_div)
        //setup all things in the new_div
        new_div.classList.add('new_div')
        // body.style.display = 'block'
        // fill the div with the icons


        for(const icon of icons){
            const clone = icon.cloneNode(true)
            clone.classList.add('icon_horizontal')
            new_div.append(clone)
            console.log(clone)
        }
        isopen = true
    }

    else{
        // completely remove the div
        document.querySelector('#main_body #sidebar_logo').style.backgroundColor = 'rgb(36, 36, 134)'
        new_div.remove()
        new_div = null;
        isopen = false;
    }

})


