export default function NavBar(){
    const styles = {
        width: "80%",
        padding: "20px",
    };
    return(
        <div style={styles} id="navBar-display">
            <button class="navBarButton" id="navBarButton-upload">FileUpload</button>
            <button class="navBarButton"  id="navBarButton-loginsign">Login/Signup</button>
        </div>
    )
}