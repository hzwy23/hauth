# jQuery.i18n.properties

## About
**jQuery.i18n.properties** is a lightweight jQuery plugin for providing internationalization to javascript from ‘.properties’ files, just like in Java Resource Bundles. It loads and parses resource bundles (.properties) based on provided language and country codes (ISO-639 and ISO-3166) or language reported by browser.

Resource bundles are ‘.properties‘ files containing locale specific key-value pairs. The use of ‘.properties‘ files for translation is specially useful when sharing i18n files between Java and Javascript projects. This plugin loads the default file (eg, Messages.properties) first and then locale specific files (Messages_pt.properties, then Messages_pt_BR.properties), so that a default value is always available when there is no translation provided. Translation keys will be available to developer as javascript variables/functions (functions, if translated value contains substitutions (eg, {0}) or as a map.

This plugin was inspired on the [Localisation assistance for jQuery from Keith Wood](http://web.archive.org/web/20140517213544/http://keith-wood.name/localisation.html).

## Latest Version

1.2.4


## Features
* Use Java standard ‘.properties‘ files for translations
* Use standard ISO-639 for language code and ISO-3166 for country code
* Sequential loading of resource bundles from base language to user-specified/browser-specified so there is always a default value for an untranslated string (eg: Messages.properties, Messages_pt.properties, Messages_pt_BR.properties)
* Use browser reported language if no language was specified
* Placeholder substitution in resource bundle strings (eg, msg_hello = Hello {0}!!)
* Suport for namespaces in keys (eg, com.company.msgs.hello = Hello!)
* Support for multi-line property values
* Resource bundle keys available as Javascript vars/functions or as a map


## History
This project was originally created by [Nuno Miguel Correia Serra Fernandes](http://nunogrilo.com)
and published on Google Code. In 2014 it has been migrated to Github which is now the [official repository](https://github.com/jquery-i18n-properties/jquery-i18n-properties).

Since then, other great contributors joined the project (see Credits section below).

It has been used in various projects, including the [WebRTC phone JSCommunicator](http://jscommunicator.org) (see the demo there to
see jquery-i18n-properties in action), some [Sakai Project](http://sakaiproject.org) tools, etc.


## Example
Take as an example the following *.properties* files:

**Messages.properties:**

```ini
# This line is ignored by the plugin
msg_hello = Hello
msg_world = World
msg_complex = Good morning {0}!
```

**Messages\_pt.properties:**

```ini
# We only provide a translation for the 'msg_hello' key
msg_hello = Bom dia
```

**Messages\_pt\_BR.properties:**

```ini
# We only provide a translation for the 'msg_hello' key
msg_hello = Olá
```

Now, suppose these files are located on the ``bundle/`` folder. One can invoke the plugin like below:

```javascript
// This will initialize the plugin 
// and show two dialog boxes: one with the text "Olá World"
// and other with the text "Good morning John!" 
jQuery.i18n.properties({
    name:'Messages', 
    path:'bundle/', 
    mode:'both',
    language:'pt_BR',
    async: true,
    callback: function() {
        // We specified mode: 'both' so translated values will be
        // available as JS vars/functions and as a map

        // Accessing a simple value through the map
        jQuery.i18n.prop('msg_hello');
        // Accessing a value with placeholders through the map
        jQuery.i18n.prop('msg_complex', 'John');

        // Accessing a simple value through a JS variable
        alert(msg_hello +' '+ msg_world);
        // Accessing a value with placeholders through a JS function
        alert(msg_complex('John'));
    }
});
```
This will initialize the plugin (loading bundle files and parsing them) and show a dialog box with the text “Olá World” and other with “Good morning John!”. The english word “World” is shown because we didn’t provide a translation for the `msg_world` key. Also notice that keys are available as a map and also as javascript variables (for simple strings) and javascript functions (for strings with placeholders for substitution).

## Asynchronous Language File Loading

Synchronous Ajax has now been deprecated and will be removed at some point in the future, so web developers need to
start thinking about writing their code as callbacks (https://xhr.spec.whatwg.org/).

With this in mind ...

If you supply the flag 'async' on the settings and set it to true, all ajax calls are executed asynchronously and the
supplied callback is called after the language files have all been downloaded and parsed. If you leave the flag off,
or set it to false, the behaviour is as before: all the files are parsed synchronously and the callback is called at the
end of the process.


## Usage


### Using the plugin
1. Include it on your ``<head>`` section:

```html
<script type="text/javascript" language="JavaScript"
  src="js/jquery.i18n.properties-min.js"></script>
```

2. Initialize the plugin (minimal usage, will use language reported by browser), and access a translated value (assuming a key named ‘org.somekey‘ exists, it will be setup as a variable you can use directly in your Javascript):

```html
<script>
	jQuery.i18n.properties({
  		name: 'Messages', 
  		callback: function(){ alert( org.somekey ); }
	});
</script>
```

### Additional requirement on Firefox
If using Firefox and a Tomcat webapp, you may get a `syntax error` in the Javascript console. The solution is to tell Tomcat the properties files should be interpreted as `text/plain`. To do this, add the following to your web.xml:
`
<mime-mapping> 
        <extension>properties</extension>
        <mime-type>text/plain</mime-type> 
</mime-mapping>
`

### Building a minified JavaScript file

1. Install the closure compiler tool:

   ``apt-get update && apt-get install closure-compiler``

2. Run it:

   ``closure-compiler --js jquery.i18n.properties.js \
                    --js_output_file jquery.i18n.properties.min.js``
   

### Options             

Option | Description | Notes
------ | ----------- | -----
**name**   | Partial name (or names) of files representing resource bundles (eg, ‘Messages’ or ['Msg1','Msg2']). Defaults to 'Messages' | Optional String or String[] |
**language** | ISO-639 Language code and, optionally, ISO-3166 country code (eg, ‘en’, ‘en_US’, ‘pt_BR’). If not specified, language reported by the browser will be used instead. | Optional String |
**path** | Path to directory that contains ‘.properties‘ files to load. | Optional String |
**mode** | Option to have resource bundle keys available as Javascript vars/functions OR as a map. The ‘map’ option is mandatory if your bundle keys contain Javascript Reserved Words. Possible options: ‘vars’ (default), ‘map’ or ‘both’. | Optional String |
**debug** | Option to turn on console debug statement. Possible options: true or false. | Optional boolean |
**cache** | Whether bundles should be cached by the browser, or forcibly reloaded on each page load. Defaults to false (i.e. forcibly reloaded). | Optional boolean |
**encoding** | The encoding to request for bundles. Property file resource bundles are specified to be in ISO-8859-1 format. Defaults to UTF-8 for backward compatibility. | Optional String |
**callback** | Callback function to be called uppon script execution completion. | Optional function() |
    
                    
## Copyright, Credits and License
Copyright © 2011 Nuno Miguel Correia Serra Fernandes (nunogrilo.com)

Special thanks to great contributors:

* [Daniel Pocock](https://github.com/dpocock)
* [mlohbihler](https://github.com/mlohbihler)
* [Guillaume Gerbaud](https://github.com/ggerbaud)
* [Adrian Fish](https://github.com/adrianfish)

Licensed under the [MIT License](LICENSE).
