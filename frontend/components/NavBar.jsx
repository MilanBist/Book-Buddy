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
            <button class="navBarButton" id="navBarButton-upload" onClick={handleUploadClick}>FileUpload</button>
            <button class="navBarButton"  id="navBarButton-loginsign">Login/Signup</button>
        </div>
    )
}