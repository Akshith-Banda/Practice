console.log('hi')
var action_url = "index/"

var link = document.getElementsByClassName('link-list')
for(var i = 0; i < link.length; i++){
    link[i].addEventListener('click', function(){
        console.log(this.dataset.id)
        var id =  this.dataset.id
        var complete_url = 'complete/'+id
        var update_url = 'update/'+id
        

        document.getElementById('completeBtn').removeAttribute('disabled')
        document.getElementById('updateBtn').removeAttribute('disabled')
        document.getElementById('completeLink').setAttribute("href", complete_url)
        document.getElementById('updateLink').setAttribute("href", update_url)
        
    })
    
    
    
}



/*

*/


