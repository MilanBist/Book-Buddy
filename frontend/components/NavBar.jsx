<<<<<<< HEAD
=======
import '../styles/button.css'
import { Link } from 'react-router-dom';

>>>>>>> fc0f258 (All commit saved)
export default function NavBar({fileUploadStatusChanger, status} ){
    const styles = {
        width: "80%",
        padding: "20px",
    };

    const handleUploadClick = ()=>{
        if (status == true){
            fileUploadStatusChanger(false)
            return
        } else{
            fileUploadStatusChanger(true)
        }
    }
    return(
        <div style={styles} id="navBar-display">
<<<<<<< HEAD
            <button class="navBarButton" id="navBarButton-upload" onClick={handleUploadClick}>FileUpload</button>
            <button class="navBarButton"  id="navBarButton-loginsign">Login/Signup</button>
=======
            <button class="button" id="navBarButton-upload" onClick={handleUploadClick}>FileUpload</button>
            <Link to='/login'><button class="button"  id="navBarButton-loginsign">Login/Signup</button></Link>
>>>>>>> fc0f258 (All commit saved)
        </div>
    )
}