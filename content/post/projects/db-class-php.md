---
title: PHP Mysqli DB Class
author: Luis Rodriguez
type: post
date: 2015-09-15T22:36:25+00:00
comments: false
categories:
  - Projects
tags:
  - luis-portfolio
  - andrew-portfolio
  - maave-portfolio
  - php
  - mysqli

---

A class written to simplify the use oh mysqli in php. Uses similar concepts to other db classes built into other frameworks. [Download the file here](https://downloads.techreanimate.com/auowdk) or look at the code below.

<!--more-->

```php
<?php
	class Database{ 
		public $host     = DB_HOST;
		public $database = DB_DATABASE;
		public $username = DB_USER;
		public $password = DB_PASSWORD;
		public $link    = 0;
		public $result  = 0;
		public $record   = array();
		public $row;
		public $errNo    = 0;
		public $error    = "";
	    
		/* Call connection method upon constructing
		 * Pass in true if you want a second connection object
		 * Example:
		 * 		$db2 = new Database(true);
		 * 		$db2->database = $config['MC_DB_Username'];
		 * 		$db2->username = $config['MC_DB_Username'];
		 * 		$db2->password = $config['MC_DB_Password'];
		 * 		$db2->host = $config['MC_DB_Host'];
		 * 		$db2->createConnection();
		 * */
		public function __construct($noconnect = false){
			if(!$noconnect){
				$this->createConnection(); 
			}
		}

		/* Connection to the database */
		public function createConnection(){ 
	        if( 0 == $this->link ) 
	            $this->link=mysqli_connect( $this->host, $this->username, $this->password );
			if( !$this->link ) 
	   			$this->stopExec( "Trying to connect.... Result: failed" );
			if( !mysqli_select_db($this->link,$this->database)) 
	            $this->stopExec( "Cannot use database ".$this->database );
		} 
	
		/* Execute query string */
	 	public function query($queryString,$echo = false,$exit = true){
	 		if($echo == true){ echo "<br>".htmlspecialchars($queryString)."<br>"; } 
			$this->result = mysqli_query($this->link, $queryString);
			$this->row = 0; 
			$this->errNo = mysqli_errno($this->link);
			$this->error = mysqli_error($this->link);
			if(!$this->result){
				$this->stopExec( "Invalid SQL String: ".$queryString , $exit );
			}else{
				if(IS_DEV){
					global $clean;
					$newlink = mysqli_connect( $this->host, $this->username, $this->password );
					mysqli_select_db($newlink,$this->database);
					mysqli_query($newlink, "INSERT INTO `ServerLog`(`Message`)VALUES('".$clean->UserInput($queryString)."')");
				}
			}
			return $this->result; /* return the resource id of query. You can either then call fetchRecord method. */
	    } 
	 
		public function fetchRecord($type = 1, $result = null){
			/* Get from specific result */
			if($result != null){
				$this->result = $result;
			}
			
			if(is_bool($this->result)){
				return "boolean given, check query";
			}
			 
			/* 1 --> for array 2 --> for object */
			if($type == 1){
				$record = mysqli_fetch_array($this->result); /*return array */
			}else{
				$record = mysqli_fetch_object($this->result); /*return object */
			}
	  
			return $record;
		}
		
		/* Stop the execution of query when there's an error */
		public function stopExec($msg,$exit){ 
			serverlog("[MySQL Error][".$this->errNo."] ".$this->error." Message: ".$msg,'crucial');
			if($exit){
				header('Location: http://'.SYS_URL_PATH.'/error.php?code=10');
			}
			
			exit();
		} 
	 
		/* Get the number of row */
		public function numRows($result = null){
			if($result != null){
				return mysqli_num_rows($result);
			}else{
				return mysqli_num_rows($this->result);
			}
			
		}
		
		/* Get the the last insert id */
		public function lastID(){
			return mysqli_insert_id($this->link);
		}
		
		public function real_escape_string($str){
			return mysqli_real_escape_string($this->link,$str);
		}
	}

```