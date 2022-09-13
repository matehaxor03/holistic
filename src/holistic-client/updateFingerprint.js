const fs = require('fs');
const path = require("path");
const crypto = require('crypto');

const updateFingerprint = async function() {

  var manifestJSONContent = fs.readFileSync("./public/manifest.json");
  var manifestObj = JSON.parse(manifestJSONContent);


  var filenames = [];

  const getFilesRecursively = (directory) => {
    const filesInDirectory = fs.readdirSync(directory);
    for (const file of filesInDirectory) {
      const absolute = path.join(directory, file);

      if(absolute.endsWith("public/manifest.json")) {
        continue;
      }
    
      if(absolute.endsWith("updateFingerprint.js")) {
        continue;
      }
    
      if(absolute.includes("/node_modules/")) {
        continue;
      }
    
      if(absolute.includes("/build/")) {
        continue;
      }
    
      if (fs.statSync(absolute).isDirectory()) {
          getFilesRecursively(absolute);
      } else {
        filenames.push(absolute);
      }
    }
  };

  function checksum(hashName, value) {
    const hashSum = crypto.createHash(hashName);
    hashSum.update(value);
    return hashSum.digest('hex');
  }
  
function checksumFile(hashName, absoluteFilename) {
  const fileBuffer = fs.readFileSync(absoluteFilename);
  return checksum(hashName, fileBuffer);
}

getFilesRecursively(__dirname);

let hashes = [];
for(let filename of filenames) {
  
  if(filename.endsWith("public/manifest.json")) {
    continue;
  }

  if(filename.endsWith("updateFingerprint.js")) {
    continue;
  }

  if(filename.includes("/node_modules/")) {
    continue;
  }

  if(filename.includes("/build/")) {
    continue;
  }

  var hash = await checksumFile("sha1", filename);
  hashes.push(hash);
}

let fingerprint = checksum("sha1", hashes.join(""));
manifestObj["fingerprint"] = fingerprint;

fs.writeFile('./public/manifest.json', JSON.stringify(manifestObj, undefined, 2), 'utf8', function(err) {
  if (err) {
    console.log('An error occured while writing JSON Object to meta.json');
    return console.log(err);
  }
});

}

updateFingerprint();