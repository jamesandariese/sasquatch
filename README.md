sasquatch(1) -- Secure Sending Qfjowirh
===

## SYNOPSIS 

    sasquatch token id_rsa.pub
    sasquatch decrypt [-i id_rsa] infile

## DESCRIPTION

* Create a random token as a session key with a colleague
* Decrypt the random token on the other side

Sasquatch uses a known ssh public key to bootstrap secure communications
over an insecure line.

## OPTIONS

`sasquatch token <keyfile.pub>`
    
 * `keyfile.pub`: The public key to encrypt for.  If not specified or `-`, read from STDIN.
 * `-q=false`:
   Quiet mode doesn't print out info about the token and prints the token followed by base64 encoded encrypted content on two lines
 * `-token-chars="abcd1234"`:
   Characters to use in generating a token
 
`sasquatch decrypt <infile>`

 * `infile`:
   Input filename.  If not specified or `-`, read from STDIN.
 * `-i="/Users/james/.ssh/id_rsa"`:
   Private key file (identity) to use for decryption

Generates a random token which fills the space available in the RSA key.

## TODO

The hash still needs to be printed on both encrypt and decrypt.

## NAMING

Sasquatch was an expansion on an abbreviation that made sense at one time.  The exact origin of the name has been lost. 

## AUTHOR

James Andariese <james@strudelline.net>

## COPYRIGHT

Copyright (c) 2015, James Andariese

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
