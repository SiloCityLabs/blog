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

[The Function object provides methods for functions. In JavaScript, every function is actually a Function object.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function).

```javascript
function func1(){
    return "hello world";
}
const funcObj = createFunction1();
console.log( funcObj() ); //outputs "hello world"
```

or

```javascript
const funcObj = function func1(){
    return "hello world";
}
console.log( funcObj() ); //outputs "hello world"
```

We don't even need to assign it. Simply pass it around, use it as a param. It could be a function triggered by a forEach loop, or a callback function triggered by a fetch request

# Callbacks

Since functions are also objects, we can pass them around. [A callback function is a function passed into another function as an argument, which is then invoked inside the outer function to complete some kind of routine or action.](https://developer.mozilla.org/en-US/docs/Glossary/Callback_function).

```javascript
function greet(name, callback) {
  console.log(`Hello, ${name}!`);
  callback(); // Execute the callback function
}

function greetMany(name, ...callbacks) { // many params
  console.log(`Hello, ${name}!`);
  callbacks.forEach(  // loop through callback array
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

Callbacks are generally used for asynchronous calls - functions that run in the background, that may take some time to return a value. This looks OK now but can get messy with many callbacks.

# Arrow syntax

This is getting too long. We need shorter functions. Arrow syntax ` => ` was introduced in ES6 in 2015 to make this faster.

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

greet("Alice", saySomethingElse, sayGoodbye);
```

# Promises

# Async

# PouchDB

Our examples will use [PouchDB](https://pouchdb.com/), a JavaScript database that works in the browser. It's a document database that stores JSON objects. Every object has a unique primary key `_id`. 1 ID = 1 JSON object. 

```javascript
const db = new PouchDB('my_database');

let helloWorld = {
  _id: '001',
  title: 'Hello World'
}

db.put( helloWorld )
.then(function (doc) {
  console.log( doc.title ); //outputs 'Hello World'
}).catch(function (err) {
  console.error('Error:', err);
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
