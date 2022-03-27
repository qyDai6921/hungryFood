# hungryFood
Software:   
go1.17.8 (Enable Go modules integration: GOPROXY=https://goproxy.cn)  
MongoDB 4.2.19  
MongoDBCompass (or other database tools)  
GoLand 2020.3.4  
Postman 9.15.2  


### Step 1:  
Create your own MongoDB database as the format below:  
	MongoAddr  = "127.0.0.1:27017"  
	MongoUser  = "root"  
	MongoPass  = "123456"  
	MongoDB   = "hungry"  
You can change MongoAddr to your local IP address and use Postman to verify your GET and POST request.

There are 2 collections you should create in your database, hungry_list and hundry_order. You can import these 2 collections under “hungryFood” folder.  
hungry_list shows the information of each chef, including the id, zip code, name, open time, close time, rating, average price per person, descriptions about his/her feature with 2 tags, and 2 menu items with 2 tags and price.  
hundry_order shows the information of each custom’s requirement, including the id, zip code, username, deliver time, address, phone number, price, tax, tip, total price, order number, and 2 menu items with menu id, 2 tags, price and count.  

### Step 2:  
GET:  
http://localhost/api/search_menus?zip_code=10001&deliver_time=1648442128  
You can change the parameter values of “zip_code” and “deliver_time”, but cannot set them as null. Here are just 2 set of data in hungry_list, you can add more chefs data. Two important parameter are “zip_code” and “deliver_time”, and I set 2 options for “zip_code”, 10001 and 22030. The open time is 1648267200 and close time is 1651377600 in unix timestamp format, which means the deliver time should be between Sat Mar 26 2022 00:00:00 GMT-0400 and Sun May 01 2022 00:00:00 GMT-0400. If you enter unmatched zip code or unmatched delivery time with correct format, you will see data is null although the massage is correct.    

Here is an example result:  
![menu](https://user-images.githubusercontent.com/91996082/160298388-9a36f99e-5b4c-4638-8b15-480b8d8937f2.PNG)


### Step 3:  
POST:  
http://localhost/api/order  
If you put the information of an order with correct JSON form in the body of request blank, you will see the below massage on the response body:  
{  
    "code": 0,  
    "msg": "success",  
    "data": null  
}  
And your order will be put on the collection hundry_order in your database.  
  
Order requirements:  
1.	chief_id cannot be null;  
2.	menu_items cannot be less than or equal to 0;  
3.	price cannot be less than or equal to 0;  
4.	deliver_time cannot be less than the current time;  
5.	address cannot be null;  
6.	zip_code should be correct (here is just 10001 and 22030, you can add more in collection hungry_list);  
  
Here is an example result:  
![order](https://user-images.githubusercontent.com/91996082/160298431-16eaf139-4ab3-4efa-a310-3c9772abd7fa.PNG)  

The example results are shown in the ../screenshot_result/ folder.
