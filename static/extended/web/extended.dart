import 'dart:html';

void main() {
  querySelector("#databutton").onClick.listen(metedata);  
  querySelector("#dataget").onClick.listen(jaladata);
}

void metedata(e){
  print("emmmm");
  
  TextInputElement data = querySelector("#dataorigin");
  
  window.alert("a meter '${data.value}'");
  
  HttpRequest.getString("/setdata?data=${data.value}");
}
void jaladata(e){
  HttpRequest.getString("/getdata").then(jaladata_response);
  
  
}


void jaladata_response(response){
  querySelector("#datadestiny").text = response;
}