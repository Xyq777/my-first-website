let modalText=document.getElementById("modalText")
let modal =document.querySelector(".modal")
let img=document.querySelector("img")
username.onfocus=()=>{

    img.setAttribute('style', 'transform: translate(-60px);')
}
username.onblur=()=>{
    img.setAttribute('style', 'transform: translate(60px);')

}
let tool=document.querySelector(".tooltip div")
password.onfocus=function (){


    tool.setAttribute('style', 'visibility: visible;')
}
password.onblur=function(){
    tool.setAttribute('style', 'visibility: hidden;')
}
let text=document.querySelector(".text")


document.querySelector(".modal-close").addEventListener("click",()=>{
    modal.setAttribute('style', 'visibility:hidden;')
})

formElem.onsubmit=async (e)=>{
    e.preventDefault()

    fetch('/regReceive',{
        method:"post",
        body:new FormData(formElem)}
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






