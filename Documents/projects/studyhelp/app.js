
const button = document.querySelector('#sidebar_logo');
let isopen = false;
let new_div = null;

const sidebar_work = ()=>{
if(!isopen){
        // make new div on the main body and do other things there
        const body = document.querySelector('#main_body #body h2');
        const icons = document.querySelectorAll("#sidebar i");
        // get the mouse click option and set the logo's background to green
        const logo = document.querySelector('#main_body #sidebar_logo');
        logo.style.backgroundColor = 'green';
        // I have got the icons then view all the icons in the screen
        // make a div at the top of the of the main body
        new_div = document.createElement('div');
        body.append(new_div)
        //setup all things in the new_div
        new_div.classList.add('new_div')
        // body.style.display = 'block'
        // fill the div with the icons


        // add icons to the div in the body
        const dataset_value = ['myday', 'preview', 'timer', 'finished', 'tasks'];
        let i = 0;
        for(const icon of icons){
            const clone = icon.cloneNode(true)
            clone.classList.add('icon_horizontal')
            clone.dataset.value = dataset_value[i]
            new_div.append(clone)
            i += 1;
        }
        let new_icons = document.querySelectorAll("#main_body #body h2 div i");
        for(const icon of new_icons){
            icon.addEventListener('click', (e)=>{
                let new_value = e.target.dataset.value
                workIcon(new_value)
    })
}
        // set the isopen to be true
        isopen = true
    }

    else{
        // completely remove the div and set new_div to null again
        document.querySelector('#main_body #sidebar_logo').style.backgroundColor = 'rgb(18, 18, 20)'
        new_div.remove()
        new_div = null;
        isopen = false;
    }
}

const workIcon = (value) =>{
    if(value === 'myday'){
        console.log('myday')
    }
    if(value === 'timer'){
        console.log(value)
    }
    if(value === 'preview'){
        console.log(value)
    }
    if(value === 'finished'){
        console.log(value)
    }
    if(value === 'tasks'){
        console.log(value)
    }

}

// work for the sidebar one
button.addEventListener('click', sidebar_work);

// work for the icons when clicked
const icons = document.querySelectorAll("#main_body #sidebar i");
for(const icon of icons){
    // access the dataset value
    // let value_ = icon.dataset.value;
    // console.log(value_)
    
    icon.addEventListener('click', (e)=>{
        let new_value = e.target.dataset.value
        workIcon(new_value)
    })
}

