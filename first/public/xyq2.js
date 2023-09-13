let modalText=document.getElementById("modalText")
let modal =document.querySelector(".modal")
let img=document.querySelector("img")
username.onfocus=()=>{

    img.setAttribute('style', 'transform: translate(-60px);')
}
username.onblur=()=>{
    img.setAttribute('style', 'transform: translate(60px);')

}


let text=document.querySelector(".text")


document.querySelector(".modal-close").addEventListener("click",()=>{
    modal.setAttribute('style', 'visibility:hidden;')
})
formElem2.onsubmit=async (e)=>{
    e.preventDefault()

    fetch('/logReceive',{
        method:"post",
        body:new FormData(formElem2)}
    )
        .then
        (response=>{
            return response.json()}
        )
        .then
        (data=>{

                modalText.innerText=data.msg
                modal.setAttribute('style', 'visibility: visible;')


            }
        )
}







