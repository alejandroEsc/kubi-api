// annotations needed for the rest gateway are getting in the way of updating this. Need
// to figure out a solution.

// Do knock a door
console.log("sending message: {knockDoor:true}")
client.getHello({knockDoor:true}, printReply);

// Do NOT knock a door
console.log("sending message: {knockDoor:false}")
client.getHello({knockDoor:false}, printReply);
