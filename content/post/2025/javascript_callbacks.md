---
title: "Callbacks suck, async is easy"
author: maave
type: post
date: 2025-08-25T00:00:00+00:00
url: /post/2025/08/25/javascript-async-is-easy/
draft: true
categories:
  - JavaScript
tags:
  - promises
  - pouchdb
---

PouchDB for Javascript requires stacking quite a bit of JS-syntax together. I haven't written JS in quite a while so this was daunting.

# Synchronous PHP example

Here is a PHP sample of accessing a database. 

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
 // function may be ommitted
const funcObj = function (){
    return "hello world";
}
console.log( funcObj() ); //outputs "hello world"
```

We don't even need to assign it. Simply pass it around, use it as a param. It could be a function triggered by a forEach loop, or a callback function triggered by a fetch request

# ForEach

Introduced: ES5, 2009

Loops through an array, calls a function for each item in the array.

```javascript
numbers=[1,2,3,4]
sum = 0;
numbers.forEach(  // loop through callback array
  function(num) {
    sum += num;
  }
);
console.log("sum:"+sum); // "sum:10"

numbers.forEach( someOtherFunction );
```

# Rest Parameters

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

Callbacks are generally used for asynchronous calls - functions that run in the background, that may take some time to return a value. 

```javascript
const db = new PouchDB('my_database_');

db.remove('001'); //clear my dev record, don't handle error
//sleep, brute force waiting until the db is cleared
function sleep(ms=1000) {
  const start = Date.now();
  while ( Date.now() - start < ms ) {};
  console.log("waited " + ms + " ms");
}
sleep();

const doSomething = function(err, doc){
  // last thing to run
  console.log(doc.title);
}

const errorHandler = function(err, response) {
  if (err) { 
    return console.log(err); 
  }
  // OK response and revision# from pouchdb
  console.log(response);
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

Let's go back to the callback example. forEach is expecting a function to call and as a param it's sending an item from the array.

```javascript
callbackArray.forEach(
  function( callback ) {
    callback(); // execute
  }
);
```

One param, one line of code, we can compact that function to `callback => callback()`

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


# Async

Introduced: ES8, 2017

[The await operator is used to wait for a Promise and get its fulfillment value. It can only be used inside an async function or at the top level of a module.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/await)

```javascript
const db = new PouchDB('my_database_'+Math.random()); // getting tired of previous records

let doc = {
  _id: '001',
  title: 'Hello World'
}

async function doStuff() { // async func runs in a new thread
  try {
    const response = await db.put(doc); // wait for promise to complete
    const gotDoc = await db.get("001");
    console.log(gotDoc.title);
  } catch (err) {
    console.log(err);
  }

}
doStuff();
console.log("start");
```

Finally our code look sane and serial. Under the hood this is still using promises.


# PouchDB

Our examples will use [PouchDB](https://pouchdb.com/), a JavaScript database that works in the browser. It's a document database that stores JSON objects. Every object has a unique primary key `_id`. 1 ID = 1 JSON object. 

```javascript
const db = new PouchDB("my_database_");
db.destroy(); //hope this works fast enough, I have leftover data

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

The brute-force serial method. Wait for the object to probably, hopefully, finish.

```javascript
db.put( doc );
sleep(1000);
db.get( doc.id );
```

Once in the database, every object gets a revision `_rev` to prevent data conflicts. To update an object you must provide `_id` AND the current `_rev`. You can't blind-update a record ... well you can with `force=true` but it's not recommended. Practically that means every put() needs a get().

`db.get( id )` returns a promise. ["A Promise is an object representing the eventual completion or failure of an asynchronous operation."](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises). It's not the DB object we want, it's a running process in object form.

```javascript
db.get( doc.id );
```
