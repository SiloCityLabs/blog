---
title: "Callbacks suck, async is easy"
author: maave
type: post
date: 2025-08-25T00:00:00+00:00
url: /post/2025/08/25/javascript-async-is-easy/
draft: false
categories:
  - JavaScript
tags:
  - pouchdb
  - promises
  - async
---

I'm using [PouchDB](https://pouchdb.com/) lately, a JavaScript database that works in the browser. PouchDB for Javascript requires callbacks and I found myself stacking quite a bit of JS-specific syntax together. I haven't written JS in quite a while so this was daunting. Lets go through all the pieces that add up to _easy_ async database calls.

# Synchronous PHP example

Here is a PHP sample of accessing a database. 

<!--more-->

```php
$couch = new CouchClient('http://localhost:5984', 'my_database');

$helloWorld = [
  '_id' => '001',
  'title' => 'Hello World'
];

$couch->storeDoc($helloWorld); // Save the document to CouchDB
$doc = $couch->getDoc('001');
echo $doc->title; // outputs 'Hello World'
echo "End";
```

Each command runs serially, one after the other. Nothing fancy. The code at the end runs at the end.

Now for the Javascript version

```javascript
const db = new PouchDB('my_database');

let helloWorld = {
  _id: '001',
  title: 'Hello World'
}

db.put( helloWorld ).then(function (response) {  // Document saved successfully
    return db.get( "001" ).then(function (doc) { // get doc
        console.log( doc.title ); // use doc
    });
}).catch(function (err) {
  console.error('Error:', err);
});

console.log("End");
```

Ew, what happened to the flow? Why do I need to nest functions in my PUT? "End" is printing before "Hello World".

In Javascript this operation is asynchronous. We're not simply waiting. We're running operations in the background, in a new thread. So what happens when our background thread finishes? It needs a function to call. Who ya gonna call? Function objects!

# Functions are objects, too

[In JavaScript, every function is actually a Function object.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function).

```javascript
function func1(){
    return "hello world";
}
const funcObj = func1;
console.log( funcObj() ); //outputs "hello world"
```

[The function keyword can be used to define a function inside an expression.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/function)

```javascript
 // anonymous function, the function name may be omitted
const funcObj = function (){
    return "hello world";
}
console.log( funcObj() ); //outputs "hello world"
```

We don't even need to assign it. Simply pass it around, use it as a param. It could be a function triggered by a forEach loop, or a callback function triggered by a fetch request

# ForEach and Rest Parameters

I'll be using these features in examples but they aren't required for PouchDB

<details>
  <summary>ForEach and Rest Parameters examples</summary>

## ForEach
Introduced: ES5, 2009

Loops through an array, calls a function for each item in the array. ForEach cannot return anything, regardless of the callback function.

```javascript
numbers=[1,2,3,4]
sum = 0;
numbers.forEach(  // loop through callback array
  function(num) {
    // we can still manipulate global vars
    sum += num;
    // any return value is ignored
  }
);
console.log("sum:"+sum); // "sum:10"

numbers.forEach( someOtherFunction );
```

## Rest Parameters

Introduced: ES6, 2015

[The rest parameter syntax allows a function to accept an indefinite number of arguments as an array.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/rest_parameters). `...` syntax.

```javascript
function addNumbers(...numbers) { // many params, condense to an array
  sum = 0;
  numbers.forEach(
    function(num) {
      sum += num;
    }
  );
  console.log("sum:"+sum);
}
addNumbers(1,2,3,4); // "sum:10"
```

</details>

# Callbacks

Since functions are also objects, we can pass them around. [A callback function is a function passed into another function as an argument, which is then invoked inside the outer function to complete some kind of routine or action.](https://developer.mozilla.org/en-US/docs/Glossary/Callback_function).

```javascript
function greet(name, callback) {
  console.log(`Hello, ${name}!`);
  callback(); // Execute the callback function
}

function greetMany(name, ...callbacks) { // many params
  console.log(`Hello, ${name}!`);
  callbacks.forEach(
    function(callback) {
      callback(); // execute
    }
  );
}

function sayGoodbye() {
  console.log("Goodbye!");
}
function saySomethingElse(){
    console.log("Nice to meet you.");
}

greet("Alice", sayGoodbye); // Pass sayGoodbye as a callback
greet("Alice", saySomethingElse, sayGoodbye);
```

Callbacks are used for asynchronous calls - functions that run in the background. For example, PouchDB database calls use callbacks.

```javascript
const db = new PouchDB('my_database_'+Math.random()); // getting tired of previous records

const doSomething = function(err, doc){
  // last thing to run
  // this could update the UI with record info
  console.log(doc.title);
}

const errorHandler = function(err, response) {
  if (err) { 
    return console.log(err); 
  }
  // OK response and revision# from pouchdb
  console.log(response);
  // get doc, then run doSomething
  db.get( response.id, doSomething);
}

let helloWorld = {
  _id: '001',
  title: 'Hello World'
}

// doc, callbackFunction
db.put( helloWorld, errorHandler );
```

This looks OK now but can get messy with many callbacks.

# Arrow syntax

Introduced: ES6, 2015

This is getting too long. We need shorter functions. [Arrow syntax ` => ` will make this faster](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions).

```javascript
(parameters) => expression
// or
(parameters) => {
    //do stuff
    return value;
}
// or no params
() => { /* do stuff */ }
// or one param
param => { /* do stuff */ }
```

If the function body is a single expression, the expression's result is implicitly returned. Both the return and the curly braces {} can be omitted. Kinda like an IF one-liner.

```javascript
( param ) => param++ //automatically returns param++
// instead of
( param ) => {
    return param++;
}
```

The allows for writing one-line functions (function-objects) which can be passed around as a callback.

```javascript
const addArrow = (a, b) => a + b; // returns a+b
const plusPlus = c => c++;        // returns c+1
```

Note that this is NOT a comparator. Remember the order "less than or equals to", the "less than" comes first in a comparator: `>=`.
```javascript
const isAGreater = (a >= b); // const is a boolean, calculated at runtime. Parens not needed.
// warning: ugly code ahead
const greaterThan = (a,b) => a>=b // function that implicitly returns a boolean

console.log( 5>=2 ); // true
console.log( greaterThan(5,2) ); // true
console.log( greaterThan(2,5) ); // false
```

Let's go back to the callback example. forEach is expecting a function to call and as a param it's sending an item from the array.

```javascript
callbackArray.forEach(
  function( callback ) {
    callback(); // execute
  }
);
```

One param, one line of code. We can shorten that function to `callback => callback()`

```javascript
callbackArray.forEach(
  cb => cb()
);
```

```javascript
function greetMany(name, ...callbacks) {
  console.log(`Hello, ${name}!`);
  callbacks.forEach( cb => cb() ); // Execute all callback functions
}

function sayGoodbye() {
  console.log("Goodbye!");
}
function saySomethingElse(){
    console.log("Nice to meet you.");
}

greetMany("Alice", saySomethingElse, sayGoodbye);
```

# PouchDB

PouchDB a document database that stores JSON objects. Every object has a unique primary key `_id`. 1 ID = 1 JSON object. Once in the database, every object gets a revision `_rev` to prevent data conflicts. To update an object you must provide `_id` AND the current `_rev`. Practically that means every put() needs a get().

PouchDB can use callbacks, promises, or async. `db.get( id )` returns a promise. ["A Promise is an object representing the eventual completion or failure of an asynchronous operation."](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises). It's not the DB object we want, it's a running process in object form.


```javascript
const db = new PouchDB("my_database_"+Math.random()); // getting tired of previous records

let helloWorld = {
  _id: "001",
  title: "Hello World"
}

db.put( helloWorld ).then(function (doc) {
  //this doc is the "ok" response from PouchDB
  // {\"ok\":true,\"id\":\"001\",\"rev\":\"1-69c3951a3eb33c5c9a3067dfbff77ece\"}
  console.log( JSON.stringify(doc) );
}).catch(function (err) {
  console.error("Error:", err);
});

console.log("End");
```

# Promises

[A Promise is an object representing the eventual completion or failure of an asynchronous operation. Essentially, a promise is a returned object to which you attach callbacks, instead of passing callbacks into a function.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises)

Straight from the [pouchDB docs we can find the promise version of this](https://pouchdb.com/api.html#delete_document)

```javascript
const db = new PouchDB('my_database_'+Math.random()); // getting tired of previous records

let doc = {
  _id: '001',
  title: 'Hello World'
}

const log = function(doc) {
  console.log(doc.title);
}

db.put(doc)
.then( function (response) {
  // OK response
  console.log( "OK\n" + JSON.stringify(response) );
  return db.get("001").then(log); // return a promise object so the errors can be caught
}).catch(function (err) {
  console.log( "error" );
  console.log( JSON.stringify(err) );
});

```

# Async & Await

Introduced: ES8, 2017

[The await operator is used to wait for a Promise and get its fulfillment value. It can only be used inside an async function or at the top level of a module.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/await)

```javascript
const db = new PouchDB('my_database_');

let doc = {
  _id: '001',
  title: 'Hello World'
}

// async func runs in a new thread
async function doStuff() {
  // wipe old records if they exist
  try {
    const gotDoc = await db.get("001");  // wait for promise to complete
    await db.remove(gotDoc);
  } catch (err) {
    console.log("Nothing to remove");
  }

  // put and get
  try {
    const response = await db.put(doc); // wait for promise to complete
    const gotDoc = await db.get("001");
    console.log(gotDoc.title);
    console.log("end");
  } catch (err) {  
    // all promise errors caught at the end
    console.log("Error");
    console.log(err);
  }

}
doStuff();
console.log("start");
```

Finally, our code look sane! PUT runs first, then GET, then we can use the object. The objects we're getting are the values we want, rather than Promise wrappers. Under the hood, this is still using promises. Mentally, it's simple serial operations. Just remember to catch errors since our promises aren't handling that anymore.

# Combined samples

## Promises and arrow functions

[In this example from BeyBuilder X](https://github.com/fabelavalon/BeyBuilderX/blob/d645bf8934b97bcf62f6c5afb71d47da177fa2bc/main.js#L750), we need to update 2 records at a time. This was a pain in the butt with full functions and PouchDB revision management, but became easy when I could compact things with arrow function and use callbacks to avoid code duplication. The final function is quite readable and the promise chain runs in logical order.

Check out the function `updateField()` and an example of using it:

```javascript
updateField(record1Id, bey => bey.wko++)

function updateField(id, updater) {
  return recordsDBX.get(id).then(doc => { // .then( anonymousFunction )
    // run the function that was passed in
    updater(doc); // doc obj is modified in-place
    return recordsDBX.put(doc); // put() the modified doc, return a promise
  });
}
```

The first param `id` is our object to update. The second param `updater` is a function to manipulate the object. updateField() will get() the doc, then run the function we passed in, then put() the modified object. This allows us to update a record with a single line of code, passing in whatever function we want to modify the object.

If we mentally unroll that it looks like this:

```javascript
// unrolled version of the anonymous function we're passing in
function anonymousUpdaterFunction(bey) {
  return bey.wko++;
}

function updateField(id, updater) {
  return recordsDBX.get(id).then(doc => {
    anonymousUpdaterFunction(doc);
    return recordsDBX.put(doc);
  });
}
```

We're modifying that `doc` or `bey` in place because JavaScript passes objects by reference. The `doc` object inside updateField() is the same object as `bey` in our anonymous function.

It also would work with passing by value, since an anonymous function like `bey => bey.wko++` implicitly returns:

```javascript
function updateField(id, updater) {
  return recordsDBX.get(record1Id).then( 
    doc => recordsDBX.put( updater(doc) ) 
  );
}
```

This is getting less readable and the order of operations is getting confusing. We can compact it one more time:

```javascript
const updateField = (id, updater) => recordsDBX.get(id).then( doc => recordsDBX.put( updater(doc) ) );
```


### Full example from BeyBuilder X
```javascript
/**
 * update a vsRecord, using whatever function you pass it
 * returns promise from pouchdb
 * example usage: increment wko (KO wins)
 * updateField(record1Id, d => d.wko++).then(updateUI);
 * @param {string} id - beyblade id
 * @param {*} updater - provide your own function
 * @returns promise
 */
function updateField(id, updater) {
  return recordsDBX.get(id).then(doc => {
    // run the function that was passed in
    updater(doc); // doc obj is being modified in-place, it's passed by ref to anonymous function "updater"
    return recordsDBX.put(doc); // puts the modified doc, returns a promise
  });
}

//update the records database with a result is chosen
function updateRecords(winner, loser, outcome){
    // ID (primary key) is the 2 IDs combined
    var record1Id = winner.id + " " + loser.id;
    // create both and update both
    var record2Id = loser.id + " " + winner.id;
    
    /*
    addRecord() was too long, for full source check the git link for BeyBuilder X
    addRecord returns a promise like: return recordsDBX.put(winRecord);
    */
    promiseChain = addRecord(winner, loser) // create if they don't exist
    .then(() => addRecord(loser, winner)) 
    .then(() => {
        console.log("update record");
        // collect promises (DB updates) for both records
        let promises = [];

        switch (outcome) {
          case "KO":
              promises.push(updateField(record1Id, bey => bey.wko++)); // add to the stack of promises
              promises.push(updateField(record2Id, bey => bey.lko++));
              break;
          case "SO":
              promises.push(updateField(record1Id, bey => bey.wso++));
              promises.push(updateField(record2Id, bey => bey.lso++));
              break;
          case "burst":
              promises.push(updateField(record1Id, bey => bey.wbst++));
              promises.push(updateField(record2Id, bey => bey.lbst++));
              break;
          case "x":
              promises.push(updateField(record1Id, bey => bey.wx++));
              promises.push(updateField(record2Id, bey => bey.lx++));
              break;
          case "draw":
              promises.push(updateField(record1Id, bey => bey.draws++));
              promises.push(updateField(record2Id, bey => bey.draws++));
              break;
          case "update":
              promises.push(updateField(record1Id, bey => {
                  bey.challenger = winner;
                  bey.defender = loser;
              }));
              promises.push(updateField(record2Id, bey => {
                  bey.challenger = loser;
                  bey.defender = winner;
              }));
              break;
          default:
              console.log("Something went wrong. Record not added");
        }
        // submit whatever results are in the list
        return Promise.all(promises);
    })
    .then(() => {
        // all DB updates finished, now update UI
        return displayRecords();
    });

    return promiseChain;
}


```