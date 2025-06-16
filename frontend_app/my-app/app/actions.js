export async function deliverMessage(message) {
  const response = await httpGetAsync("http://localhost:8080/order/"+message);
  return response
}

function httpGetAsync(theUrl, callback)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() { 
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(xmlHttp.responseText);
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous 
    xmlHttp.setRequestHeader("Access-Control-Allow-Headers", "*");
    xmlHttp.send(null);
}
