package colornames

var testdata = `<!--
  Scalable Vector Graphics (SVG) 1.1 (Second Edition)
  Chapter 4: Basic Data Types and Interfaces

  $Id: types.html,v 1.1 2011/08/10 03:35:33 schepers Exp $

  Note: This document is generated from ../master/types.html.
  Run "make" from ../master/ to regenerate it.
  -->

<!DOCTYPE html
  PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"><head><meta http-equiv="Content-Type" content="text/html;charset=UTF-8"/><title>Basic Data Types and Interfaces – SVG 1.1 (Second Edition)</title><style type="text/css">
    .aliceblue { background: rgb(240, 248, 255) }
    .antiquewhite { background: rgb(250, 235, 215) }
    .aqua { background: rgb( 0, 255, 255); }
    .aquamarine { background: rgb(127, 255, 212) }
    .azure { background: rgb(240, 255, 255) }
    .beige { background: rgb(245, 245, 220) }
    .bisque { background: rgb(255, 228, 196) }
    .black { background: rgb( 0, 0, 0) }
    .blanchedalmond { background: rgb(255, 235, 205) }
    .blue { background: rgb( 0, 0, 255) }
    .blueviolet { background: rgb(138, 43, 226) }
    .brown { background: rgb(165, 42, 42) }
    .burlywood { background: rgb(222, 184, 135) }
    .cadetblue { background: rgb( 95, 158, 160) }
    .chartreuse { background: rgb(127, 255, 0) }
    .chocolate { background: rgb(210, 105, 30) }
    .coral { background: rgb(255, 127, 80) }
    .cornflowerblue { background: rgb(100, 149, 237) }
    .cornsilk { background: rgb(255, 248, 220) }
    .crimson { background: rgb(220, 20, 60) }
    .cyan { background: rgb( 0, 255, 255) }
    .darkblue { background: rgb( 0, 0, 139) }
    .darkcyan { background: rgb( 0, 139, 139) }
    .darkgoldenrod { background: rgb(184, 134, 11) }
    .darkgray { background: rgb(169, 169, 169) }
    .darkgreen { background: rgb( 0, 100, 0) }
    .darkgrey { background: rgb(169, 169, 169) }
    .darkkhaki { background: rgb(189, 183, 107) }
    .darkmagenta { background: rgb(139, 0, 139) }
    .darkolivegreen { background: rgb( 85, 107, 47) }
    .darkorange { background: rgb(255, 140, 0) }
    .darkorchid { background: rgb(153, 50, 204) }
    .darkred { background: rgb(139, 0, 0) }
    .darksalmon { background: rgb(233, 150, 122) }
    .darkseagreen { background: rgb(143, 188, 143) }
    .darkslateblue { background: rgb( 72, 61, 139) }
    .darkslategray { background: rgb( 47, 79, 79) }
    .darkslategrey { background: rgb( 47, 79, 79) }
    .darkturquoise { background: rgb( 0, 206, 209) }
    .darkviolet { background: rgb(148, 0, 211) }
    .deeppink { background: rgb(255, 20, 147) }
    .deepskyblue { background: rgb( 0, 191, 255) }
    .dimgray { background: rgb(105, 105, 105) }
    .dimgrey { background: rgb(105, 105, 105) }
    .dodgerblue { background: rgb( 30, 144, 255) }
    .firebrick { background: rgb(178, 34, 34) }
    .floralwhite { background: rgb(255, 250, 240) }
    .forestgreen { background: rgb( 34, 139, 34) }
    .fuchsia { background: rgb(255, 0, 255) }
    .gainsboro { background: rgb(220, 220, 220) }
    .ghostwhite { background: rgb(248, 248, 255) }
    .gold { background: rgb(255, 215, 0) }
    .goldenrod { background: rgb(218, 165, 32) }
    .gray { background: rgb(128, 128, 128) }
    .grey { background: rgb(128, 128, 128) }
    .green { background: rgb( 0, 128, 0) }
    .greenyellow { background: rgb(173, 255, 47) }
    .honeydew { background: rgb(240, 255, 240) }
    .hotpink { background: rgb(255, 105, 180) }
    .indianred { background: rgb(205, 92, 92) }
    .indigo { background: rgb( 75, 0, 130) }
    .ivory { background: rgb(255, 255, 240) }
    .khaki { background: rgb(240, 230, 140) }
    .lavender { background: rgb(230, 230, 250) }
    .lavenderblush { background: rgb(255, 240, 245) }
    .lawngreen { background: rgb(124, 252, 0) }
    .lemonchiffon { background: rgb(255, 250, 205) }
    .lightblue { background: rgb(173, 216, 230) }
    .lightcoral { background: rgb(240, 128, 128) }
    .lightcyan { background: rgb(224, 255, 255) }
    .lightgoldenrodyellow { background: rgb(250, 250, 210) }
    .lightgray { background: rgb(211, 211, 211) }
    .lightgreen { background: rgb(144, 238, 144) }
    .lightgrey { background: rgb(211, 211, 211) }
    .lightpink { background: rgb(255, 182, 193) }
    .lightsalmon { background: rgb(255, 160, 122) }
    .lightseagreen { background: rgb( 32, 178, 170) }
    .lightskyblue { background: rgb(135, 206, 250) }
    .lightslategray { background: rgb(119, 136, 153) }
    .lightslategrey { background: rgb(119, 136, 153) }
    .lightsteelblue { background: rgb(176, 196, 222) }
    .lightyellow { background: rgb(255, 255, 224) }
    .lime { background: rgb( 0, 255, 0) }
    .limegreen { background: rgb( 50, 205, 50) }
    .linen { background: rgb(250, 240, 230) }
    .magenta { background: rgb(255, 0, 255) }
    .maroon { background: rgb(128, 0, 0) }
    .mediumaquamarine { background: rgb(102, 205, 170) }
    .mediumblue { background: rgb( 0, 0, 205) }
    .mediumorchid { background: rgb(186, 85, 211) }
    .mediumpurple { background: rgb(147, 112, 219) }
    .mediumseagreen { background: rgb( 60, 179, 113) }
    .mediumslateblue { background: rgb(123, 104, 238) }
    .mediumspringgreen { background: rgb( 0, 250, 154) }
    .mediumturquoise { background: rgb( 72, 209, 204) }
    .mediumvioletred { background: rgb(199, 21, 133) }
    .midnightblue { background: rgb( 25, 25, 112) }
    .mintcream { background: rgb(245, 255, 250) }
    .mistyrose { background: rgb(255, 228, 225) }
    .moccasin { background: rgb(255, 228, 181) }
    .navajowhite { background: rgb(255, 222, 173) }
    .navy { background: rgb( 0, 0, 128) }
    .oldlace { background: rgb(253, 245, 230) }
    .olive { background: rgb(128, 128, 0) }
    .olivedrab { background: rgb(107, 142, 35) }
    .orange { background: rgb(255, 165, 0) }
    .orangered { background: rgb(255, 69, 0) }
    .orchid { background: rgb(218, 112, 214) }
    .palegoldenrod { background: rgb(238, 232, 170) }
    .palegreen { background: rgb(152, 251, 152) }
    .paleturquoise { background: rgb(175, 238, 238) }
    .palevioletred { background: rgb(219, 112, 147) }
    .papayawhip { background: rgb(255, 239, 213) }
    .peachpuff { background: rgb(255, 218, 185) }
    .peru { background: rgb(205, 133, 63) }
    .pink { background: rgb(255, 192, 203) }
    .plum { background: rgb(221, 160, 221) }
    .powderblue { background: rgb(176, 224, 230) }
    .purple { background: rgb(128, 0, 128) }
    .red { background: rgb(255, 0, 0) }
    .rosybrown { background: rgb(188, 143, 143) }
    .royalblue { background: rgb( 65, 105, 225) }
    .saddlebrown { background: rgb(139, 69, 19) }
    .salmon { background: rgb(250, 128, 114) }
    .sandybrown { background: rgb(244, 164, 96) }
    .seagreen { background: rgb( 46, 139, 87) }
    .seashell { background: rgb(255, 245, 238) }
    .sienna { background: rgb(160, 82, 45) }
    .silver { background: rgb(192, 192, 192) }
    .skyblue { background: rgb(135, 206, 235) }
    .slateblue { background: rgb(106, 90, 205) }
    .slategray { background: rgb(112, 128, 144) }
    .slategrey { background: rgb(112, 128, 144) }
    .snow { background: rgb(255, 250, 250) }
    .springgreen { background: rgb( 0, 255, 127) }
    .steelblue { background: rgb( 70, 130, 180) }
    .tan { background: rgb(210, 180, 140) }
    .teal { background: rgb( 0, 128, 128) }
    .thistle { background: rgb(216, 191, 216) }
    .tomato { background: rgb(255, 99, 71) }
    .turquoise { background: rgb( 64, 224, 208) }
    .violet { background: rgb(238, 130, 238) }
    .wheat { background: rgb(245, 222, 179) }
    .white { background: rgb(255, 255, 255) }
    .whitesmoke { background: rgb(245, 245, 245) }
    .yellow { background: rgb(255, 255, 0) }
    .yellowgreen { background: rgb(154, 205, 50) }
  </style><link rel="stylesheet" type="text/css" media="screen" href="style/svg-style.css"/><link rel="stylesheet" href="http://www.w3.org/StyleSheets/TR/W3C-REC" type="text/css" media="screen"/></head><body><div class="header top"><span class="namedate">SVG 1.1 (Second Edition) – 16 August 2011</span><a href="Overview.html">Top</a> ⋅ <a href="expanded-toc.html">Contents</a> ⋅ <a href="render.html">Previous</a> ⋅ <a href="struct.html">Next</a> ⋅ <a href="eltindex.html">Elements</a> ⋅ <a href="attindex.html">Attributes</a> ⋅ <a href="propidx.html">Properties</a></div>

<h1>4 Basic Data Types and Interfaces</h1><h2 id="toc">Contents</h2><ul class="toc">
<li><a href="#syntax">4.1 Syntax</a></li>
<li><a href="#BasicDataTypes">4.2 Basic data types</a></li>
<li><a href="#Precision">4.3 Real number precision</a></li>
<li><a href="#ColorKeywords">4.4 Recognized color keyword names</a></li>
<li><a href="#BasicDOMInterfaces">4.5 Basic DOM interfaces</a><ul class="toc">
<li><a href="#InterfaceSVGElement">4.5.1 Interface SVGElement</a></li>
<li><a href="#InterfaceSVGAnimatedBoolean">4.5.2 Interface SVGAnimatedBoolean</a></li>
<li><a href="#InterfaceSVGAnimatedString">4.5.3 Interface SVGAnimatedString</a></li>
<li><a href="#InterfaceSVGStringList">4.5.4 Interface SVGStringList</a></li>
<li><a href="#InterfaceSVGAnimatedEnumeration">4.5.5 Interface SVGAnimatedEnumeration</a></li>
<li><a href="#InterfaceSVGAnimatedInteger">4.5.6 Interface SVGAnimatedInteger</a></li>
<li><a href="#InterfaceSVGNumber">4.5.7 Interface SVGNumber</a></li>
<li><a href="#InterfaceSVGAnimatedNumber">4.5.8 Interface SVGAnimatedNumber</a></li>
<li><a href="#InterfaceSVGNumberList">4.5.9 Interface SVGNumberList</a></li>
<li><a href="#InterfaceSVGAnimatedNumberList">4.5.10 Interface SVGAnimatedNumberList</a></li>
<li><a href="#InterfaceSVGLength">4.5.11 Interface SVGLength</a></li>
<li><a href="#InterfaceSVGAnimatedLength">4.5.12 Interface SVGAnimatedLength</a></li>
<li><a href="#InterfaceSVGLengthList">4.5.13 Interface SVGLengthList</a></li>
<li><a href="#InterfaceSVGAnimatedLengthList">4.5.14 Interface SVGAnimatedLengthList</a></li>
<li><a href="#InterfaceSVGAngle">4.5.15 Interface SVGAngle</a></li>
<li><a href="#InterfaceSVGAnimatedAngle">4.5.16 Interface SVGAnimatedAngle</a></li>
<li><a href="#InterfaceSVGColor">4.5.17 Interface SVGColor</a></li>
<li><a href="#InterfaceSVGICCColor">4.5.18 Interface SVGICCColor</a></li>
<li><a href="#InterfaceSVGRect">4.5.19 Interface SVGRect</a></li>
<li><a href="#InterfaceSVGAnimatedRect">4.5.20 Interface SVGAnimatedRect</a></li>
<li><a href="#InterfaceSVGUnitTypes">4.5.21 Interface SVGUnitTypes</a></li>
<li><a href="#InterfaceSVGStylable">4.5.22 Interface SVGStylable</a></li>
<li><a href="#InterfaceSVGLocatable">4.5.23 Interface SVGLocatable</a></li>
<li><a href="#InterfaceSVGTransformable">4.5.24 Interface SVGTransformable</a></li>
<li><a href="#InterfaceSVGTests">4.5.25 Interface SVGTests</a></li>
<li><a href="#InterfaceSVGLangSpace">4.5.26 Interface SVGLangSpace</a></li>
<li><a href="#InterfaceSVGExternalResourcesRequired">4.5.27 Interface SVGExternalResourcesRequired</a></li>
<li><a href="#InterfaceSVGFitToViewBox">4.5.28 Interface SVGFitToViewBox</a></li>
<li><a href="#InterfaceSVGZoomAndPan">4.5.29 Interface SVGZoomAndPan</a></li>
<li><a href="#InterfaceSVGViewSpec">4.5.30 Interface SVGViewSpec</a></li>
<li><a href="#InterfaceSVGURIReference">4.5.31 Interface SVGURIReference</a></li>
<li><a href="#InterfaceSVGCSSRule">4.5.32 Interface SVGCSSRule</a></li>
<li><a href="#InterfaceSVGRenderingIntent">4.5.33 Interface SVGRenderingIntent</a></li></ul></li></ul>

<h2 id="syntax">4.1 Syntax</h2>

  <p>The EBNF grammar is as used in the <a href="http://www.w3.org/TR/REC-xml/#sec-notation">XML specification</a>,
  with the addition of <strong>~</strong>, a <em>case-insensitive literal</em>: characters in the ASCII range (only) are declared to be case-insensitive. For example, ~"Hello" will match (H|h)(e|e)(l|L)(l|L)(o|O). This makes the productions much easier to read.</p>
  <table><tr><td>?</td><td>optional, zero or one</td></tr><tr><td>+</td><td>one or more</td></tr><tr><td>*</td><td>zero or more</td></tr><tr><td>|</td><td>alternation</td></tr><tr><td>"string"</td><td>literal</td></tr><tr><td>~"string"</td><td>case-insensitive literal</td></tr><tr><td>[]</td><td>a character range</td></tr><tr><td>[^]</td><td>excluded character range</td></tr><tr><td>()</td><td>grouping</td></tr></table>

<h2 id="BasicDataTypes">4.2 Basic data types</h2>

<p>This section defines a number of common data types used in the definitions
of SVG properties and attributes.  Some data types that are not referenced by
multiple properties and attributes are defined inline in subsequent chapters.</p>
  
  <p>Note that, as noted below, the specification of some types is different for CSS property values in style sheets (in the <a href="styling.html#StyleAttribute"><span class="attr-name">‘style’</span></a> attribute,
    <a href="styling.html#StyleElement"><span class="element-name">‘style’</span></a> element or an external style sheet) than it is for for XML attribute values (including <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a>s). This is due to restrictions in the CSS grammar. For example, scientific notation is allowed in attributes, including presentation attributes, but not in style sheets.</p>

<dl class="definitions"><dt id="DataTypeAngle"><span class="SVG-Term">&lt;angle&gt;</span></dt><dd>
    <p>Angles are specified in one of two ways depending upon whether
    they are used in CSS <a href="intro.html#TermProperty"><span class="svg-term">property</span></a> syntax or SVG <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a>
    syntax:</p>
    <ul><li>
	<p>When an &lt;angle&gt; is used in a style sheet or with a
	<a href="intro.html#TermProperty"><span class="svg-term">property</span></a> in a <a href="styling.html#StyleAttribute"><span class="attr-name">‘style’</span></a> attribute, the
	syntax must match the following pattern:</p>
	<pre class="grammar"><span id="Angle">angle</span> ::= <a href="#DataTypeNumber">number</a> (~"deg" | ~"grad" | ~"rad")?</pre>
	<p>where <span class="prop-value">deg</span> indicates degrees,
	<span class="prop-value">grad</span> indicates grads and
	<span class="prop-value">rad</span> indicates radians.
	The unit identifier may be in lower (recommended) or upper case.</p>
	<p>For properties defined in CSS2
	[<a href="http://www.w3.org/TR/2008/REC-CSS2-20080411/">CSS2</a>],
	an angle unit identifier must be provided (for non-zero
	values).  For SVG-specific <a href="intro.html#TermProperty"><span class="svg-term">properties</span></a> the angle unit
	identifier is optional.  If a unit is not provided, the angle
	value is assumed to be in degrees.</p>
      </li><li>
	<p>When an &lt;angle&gt; is used in an SVG <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a>,
	the syntax must match the following pattern:</p>
	<pre class="grammar">angle ::= <a href="#DataTypeNumber">number</a> ("deg" | "grad" | "rad")?</pre>
	<p>The unit identifier, if present, must be in lower case; if
	not present, the angle value is assumed to be in degrees.</p>
      </li></ul>
    <p>In the SVG DOM, &lt;angle&gt; values are represented using
      <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> or <a class="idlinterface" href="types.html#InterfaceSVGAnimatedAngle">SVGAnimatedAngle</a> objects.</p>
  </dd><dt id="DataTypeAnything"><span class="SVG-Term">&lt;anything&gt;</span></dt><dd>
    <p>The basic type &lt;anything&gt; is a sequence of zero or more characters.
    Specifically:</p>
    <pre class="grammar"><span id="Anything">anything</span> ::= <a href="http://www.w3.org/TR/2008/REC-xml-20081126/#NT-Char">Char</a>*</pre>
    <p>where <a href="http://www.w3.org/TR/2008/REC-xml-20081126/#NT-Char">Char</a> is the
    production for a character, as defined in XML 1.0
    ([<a href="refs.html#ref-XML10">XML10</a>], section 2.2).</p>
  </dd><dt id="DataTypeColor"><span class="SVG-Term">&lt;color&gt;</span></dt><dd>
    <p>The basic type &lt;color&gt; is a CSS2 compatible specification for a
      color in the sRGB color space [<a href="refs.html#ref-SRGB">SRGB</a>].
      &lt;color&gt; applies to SVG's use of the <a href="color.html#ColorProperty"><span class="prop-name">‘color’</span></a> property and
      is a component of the definitions of properties <a href="painting.html#FillProperty"><span class="prop-name">‘fill’</span></a>,
      <a href="painting.html#StrokeProperty"><span class="prop-name">‘stroke’</span></a>, <a href="pservers.html#StopColorProperty"><span class="prop-name">‘stop-color’</span></a>, <a href="filters.html#FloodColorProperty"><span class="prop-name">‘flood-color’</span></a> and
      <a href="filters.html#LightingColorProperty"><span class="prop-name">‘lighting-color’</span></a>, which also offer optional ICC-based color
      specifications.</p>
    <p>SVG supports all of the syntax alternatives for &lt;color&gt; defined in
      <a href="http://www.w3.org/TR/2008/REC-CSS2-20080411/syndata.html#value-def-color">CSS2
        syntax and basic data types</a> ([<a href="refs.html#ref-CSS2">CSS2</a>],
      section 4.3.6), with the exception that SVG allows an expanded list of
      <a href="types.html#ColorKeywords">recognized color keywords names</a>.</p>
    <p>A &lt;color&gt; is either a keyword (see
      <a href="types.html#ColorKeywords">Recognized color keyword names</a>) or a
      numerical RGB specification.</p>
    <p>In addition to these color keywords, users may specify
      keywords that correspond to the colors used by objects in the
      user's environment. The normative definition of these
      keywords is found in
      <a href="http://www.w3.org/TR/2008/REC-CSS2-20080411/ui.html#system-colors">User preferences
        for colors</a> ([<a href="refs.html#ref-CSS2">CSS2</a>], section 18.2).</p>
    <p>The format of an RGB value in hexadecimal notation is a "#"
      immediately followed by either three or six hexadecimal
      characters. The three-digit RGB notation (#rgb) is converted
      into six-digit form (#rrggbb) by replicating digits, not by
      adding zeros. For example, <span class="attr-value">#fb0</span>
      expands to <span class="attr-value">#ffbb00</span>. This
      ensures that white (<span class="attr-value">#ffffff</span>) can be
      specified with the short notation (<span class="attr-value">#fff</span>)
      and removes any dependencies on the color depth of the display. 
      
      The format of an RGB value in the
      functional notation is an RGB start-function  followed by a comma-separated
      list of three numerical values (either three integer values
      or three percentage values) followed by ")". 
      An RGB start-function is the case-insensitive string "rgb(", for example "RGB(" or "rGb(".
      For compatibility, the all-lowercase form "rgb(" is preferred.
      
      The integer
      value 255 corresponds to 100%, and to F or FF in the
      hexadecimal notation: <span class="attr-value">rgb(255,255,255)</span>
      = <span class="attr-value">rgb(100%,100%,100%)</span>
      = <span class="attr-value">#FFF</span>. White space characters are allowed
      around the numerical values. All RGB colors are specified in the sRGB
      color space [<a href="refs.html#ref-SRGB">SRGB</a>]. Using
      sRGB provides an unambiguous and objectively measurable definition
      of the color, which can be related to international standards
      (see [<a href="refs.html#ref-COLORIMETRY">COLORIMETRY</a>]).</p>
    <pre class="grammar"><span id="Color">color</span>    ::= "#" <a href="#HexDigit">hexdigit</a> <a href="#HexDigit">hexdigit</a> <a href="#HexDigit">hexdigit</a> (<a href="#HexDigit">hexdigit</a> <a href="#HexDigit">hexdigit</a> <a href="#HexDigit">hexdigit</a>)?
             | "rgb(" <a href="#WSP">wsp</a>* <a href="#Integer">integer</a> <a href="#Comma">comma</a> <a href="#Integer">integer</a> <a href="#Comma">comma</a> <a href="#Integer">integer</a> <a href="#WSP">wsp</a>* ")"
             | "rgb(" <a href="#WSP">wsp</a>* <a href="#Integer">integer</a> "%" <a href="#Comma">comma</a> <a href="#Integer">integer</a> "%" <a href="#Comma">comma</a> <a href="#Integer">integer</a> "%" <a href="#WSP">wsp</a>* ")"
             | <var>color-keyword</var>
<span id="HexDigit">hexdigit</span> ::= [0-9A-Fa-f]
<span id="Comma">comma</span>    ::= <a href="#WSP">wsp</a>* "," <a href="#WSP">wsp</a>*
    </pre>
    <p>where <var>color-keyword</var> matches (case insensitively) one of the
      color keywords listed in <a href="#ColorKeywords">Recognized color keyword names</a>
      below, or one of the system color keywords listed in
      <a href="http://www.w3.org/TR/2008/REC-CSS2-20080411/ui.html#system-colors">User preferences
        for colors</a> ([<a href="refs.html#ref-CSS2">CSS2</a>], section 18.2).</p>
    <p>The corresponding SVG DOM interface definitions for
      &lt;color&gt; are defined in
      <a href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html">Document Object Model CSS</a>;
      in particular, see <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-RGBColor">RGBColor</a>
      ([<a href="refs.html#ref-DOM2STYLE">DOM2STYLE</a>], section 2.2).
      SVG's extension to color, including the ability to specify ICC-based colors,
      are represented using DOM interface <a class="idlinterface" href="types.html#InterfaceSVGColor">SVGColor</a>.</p>
  </dd><dt id="DataTypeCoordinate"><span class="SVG-Term">&lt;coordinate&gt;</span></dt><dd>
    <p>A &lt;coordinate&gt; is a length in the user coordinate system that is
      the given distance from the origin of the user coordinate system along the
      relevant axis (the x-axis for X coordinates, the y-axis for Y
      coordinates).  Its syntax is the same as that for
      <a href="#DataTypeLength">&lt;length&gt;</a>.</p>
    <pre>coordinate ::= <a href="#Length">length</a></pre>
    <p>Within the SVG DOM, a &lt;coordinate&gt; is represented as
      an <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> or an <a class="idlinterface" href="types.html#InterfaceSVGAnimatedLength">SVGAnimatedLength</a>.</p>
  </dd><dt id="DataTypeFrequency"><span class="SVG-Term">&lt;frequency&gt;</span></dt><dd>
    <p>Frequency values are used with aural properties. As defined in
      CSS2, a frequency value is a &lt;number&gt; immediately followed
      by a frequency unit identifier.  The frequency unit identifiers
      are:</p>
    <ul><li><span class="attr-value">Hz</span>: Hertz</li><li><span class="attr-value">kHz</span>: kilo Hertz</li></ul>
    <p>Frequency values may not be negative.</p>
    <p>In the SVG DOM, &lt;frequency&gt; values are represented
      using the <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSPrimitiveValue">CSSPrimitiveValue</a>
      interface defined in <a href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html">Document Object Model CSS</a>
      ([<a href="refs.html#ref-DOM2STYLE">DOM2STYLE</a>], section 2.2).</p>
  </dd><dt id="DataTypeFuncIRI"><span class="SVG-Term">&lt;FuncIRI&gt;</span></dt><dd>Functional notation for an <a href="intro.html#TermIRIReference"><span class="svg-term">IRI</span></a>: "url(" <a href="#DataTypeIRI">&lt;IRI&gt;</a> ")".
  </dd><dt id="DataTypeICCColor"><span class="SVG-Term">&lt;icccolor&gt;</span></dt><dd>
    <p>An &lt;icccolor&gt; is an ICC color specification.  In SVG 1.1,
    an ICC color specification is given by a name, which references
    a <a href="color.html#ColorProfileElement"><span class="element-name">‘color-profile’</span></a> element, and one or more color component
    values.  The grammar is as follows:</p>
    <pre class="grammar"><span id="ICCColor">icccolor</span> ::= "icc-color(" <a href="#Name">name</a> (<a href="#CommaWSP">comma-wsp</a> <a href="#Number">number</a>)+ ")"
name     ::= [^,()#x20#x9#xD#xA] /* any char except ",", "(", ")" or wsp */</pre>
    <p>The corresponding SVG DOM interface for &lt;icccolor&gt; is
    <a class="idlinterface" href="types.html#InterfaceSVGICCColor">SVGICCColor</a>.</p>
  </dd><dt id="DataTypeInteger"><span class="SVG-Term">&lt;integer&gt;</span></dt><dd>
    <p>An &lt;integer&gt; is specified as an optional sign character ("+" or
    "-") followed by one or more digits "0" to "9":</p>
    <pre class="grammar"><span id="Integer">integer</span> ::= [+-]? [0-9]+</pre>
    <p>If the sign character is not present, the number is non-negative.</p>
    <p>Unless stated otherwise for a particular attribute or
    <a href="intro.html#TermProperty"><span class="svg-term">property</span></a>, the range for an &lt;integer&gt; encompasses (at a
    minimum) -2147483648 to 2147483647.</p>
    <p>Within the SVG DOM, an &lt;integer&gt; is represented as a
    <span class="DOMInterfaceName">long</span> or an
    <a class="idlinterface" href="types.html#InterfaceSVGAnimatedInteger">SVGAnimatedInteger</a>.</p>
  </dd><dt id="DataTypeIRI"><span class="SVG-Term">&lt;IRI&gt;</span></dt><dd>
    <p>
      An Internationalized Resource Identifier (see <a href="intro.html#TermIRIReference"><span class="svg-term">IRI</span></a>).
      For the specification of <a href="intro.html#TermIRIReference"><span class="svg-term">IRI</span></a> references in SVG, see
      <a href="linking.html#IRIReference">IRI references</a>.
    </p>
  </dd><dt id="DataTypeLength"><span class="SVG-Term">&lt;length&gt;</span></dt><dd>

    <p>A length is a distance measurement, given as a number along with a unit
    which may be optional. Lengths are specified in one of two ways depending
    upon whether they are used in CSS <a href="intro.html#TermProperty"><span class="svg-term">property</span></a> syntax or SVG
    <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a> syntax:</p>
    <ul><li>
      <p>When a &lt;length&gt; is used in a style sheet or with a
      <a href="intro.html#TermProperty"><span class="svg-term">property</span></a> in a <a href="styling.html#StyleAttribute"><span class="attr-name">‘style’</span></a> attribute, the
      syntax must match the following pattern:</p>
      <pre class="grammar"><span id="Length">length</span> ::= <a href="#DataTypeNumber">number</a> (~"em" | ~"ex" | ~"px" | ~"in" | ~"cm" | ~"mm" | ~"pt" | ~"pc")?</pre>
      <p><a href="http://www.w3.org/TR/2008/REC-CSS2-20080411/syndata.html#length-units">See
      the CSS2 specification</a> for the meanings of the unit
      identifiers. The unit identifier may be in lower (recommended)
      or upper case.</p>

      <p>For properties defined in CSS2
      [<a href="http://www.w3.org/TR/2008/REC-CSS2-20080411/">CSS2</a>],
      a length unit identifier must be provided (for non-zero values).
      For SVG-specific <a href="intro.html#TermProperty"><span class="svg-term">properties</span></a>, the length unit identifier
      is optional.  If a unit is not provided, the length value
      represents a distance in the current user coordinate system.</p>
     </li><li>
      <p>When a &lt;length&gt; is used in an SVG <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a>,
      the syntax must match the following pattern:</p>
      <pre class="grammar">length ::= <a href="#DataTypeNumber">number</a> ("em" | "ex" | "px" | "in" | "cm" | "mm" | "pt" | "pc" | "%")?</pre>
      <p>The unit identifier, if present, must be in lower case; if
      not present, the length value represents a distance in the
      current user coordinate system.</p>

      <p>Note that the non-property &lt;length&gt; definition also
      allows a percentage unit identifier.  The meaning of a
      percentage length value depends on the attribute for which the
      percentage length value has been specified. Two common cases
      are: (a) when a percentage length value represents a percentage
      of the viewport width or height (refer to <a href="coords.html#Units">the section that discusses units in
      general</a>), and (b) when a percentage length value represents
      a percentage of the bounding box width or height on a given
      object (refer to <a href="coords.html#ObjectBoundingBox">the
      section that describes object bounding box units</a>).</p>

     </li></ul>

    <p>In the SVG DOM, &lt;length&gt; values are represented using
      <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> or <a class="idlinterface" href="types.html#InterfaceSVGAnimatedLength">SVGAnimatedLength</a> objects.</p>
  </dd><dt id="DataTypeListOfFamilyNames"><span class="SVG-Term">&lt;list-of-family-names&gt;</span></dt><dd>
    <p>
      A &lt;list-of-family-names&gt; is a list of font family names using the same syntax
      as the <a href="text.html#FontFamilyProperty"><span class="prop-name">‘font-family’</span></a>
      property, excluding the &lt;generic-family&gt; and <span class="attr-value">'inherit'</span>
      values.
    </p>
  </dd><dt id="DataTypeListOfStrings"><span class="SVG-Term">&lt;list-of-strings&gt;</span></dt><dd>
    <p>
      A &lt;list-of-strings&gt; consists of a separated sequence of &lt;string&gt;s.
      String lists are white space-separated, where white space is defined as one or more of the
      following consecutive characters: "space" (U+0020), "tab" (U+0009), "line feed" (U+000A) and
      "carriage return" (U+000D).
    </p>
    <p>
      The following is an EBNF grammar describing the &lt;list-of-strings&gt;
      syntax:
    </p>
<pre>
list-of-strings ::= string
                    | string wsp list-of-strings
string          ::= [^#x9#xA#xD#x20]*
wsp             ::= [#x9#xA#xD#x20]+
</pre>
  </dd><dt id="DataTypeList">
    <span id="DataTypeListOfIRI">
      <span id="DataTypeNumbers">
        <span id="DataTypeCoordinates">
          <span id="DataTypeLengths">
            <span class="SVG-Term">&lt;list-of-<var>T</var>s&gt;</span>
          </span>
        </span>  
      </span>
    </span>
  </dt><dd>
    <p>(Where <var>T</var> is a type other than
      &lt;string&gt; and
      &lt;family-name&gt;.)
       A list consists of a
    separated sequence of values. Unless explicitly described
    differently, lists within SVG's XML attributes can be either
    comma-separated, with optional white space before or after the comma,
    or white space-separated.</p>
    <p>White space in lists is defined as one or more of the
    following consecutive characters: "space" (U+0020), "tab" (U+0009),
    "line feed" (U+000A), "carriage return" (U+000D)
    and "form-feed" (U+000C).</p>
    <p>The following is a template for an EBNF grammar describing the
    &lt;list-of-<var>T</var>s&gt; syntax:</p>
    <pre class="grammar"><span id="ListOfTs">list-of-<var>T</var>s</span> ::= <var>T</var>
               | <var>T</var> <a href="#CommaWSP">comma-wsp</a> <a href="#ListOfTs">list-of-<var>T</var>s</a>
<span id="CommaWSP">comma-wsp</span>  ::= (<a href="#WSP">wsp</a>+ ","? <a href="#WSP">wsp</a>*) | ("," <a href="#WSP">wsp</a>*)
<span id="WSP">wsp</span>        ::= (#x20 | #x9 | #xD | #xA)</pre>
    <p>Within the SVG DOM, values of a &lt;list-of-<var>T</var>s&gt;
    type are represented by an interface specific for the
    particular type <var>T</var>.  For example, a &lt;list-of-lengths&gt;
    is represented in the SVG DOM using an <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a>
    or <a class="idlinterface" href="types.html#InterfaceSVGAnimatedLengthList">SVGAnimatedLengthList</a> object.</p>
  </dd><dt id="DataTypeName"><span class="SVG-Term">&lt;name&gt;</span></dt><dd>
    <p>A name, which is a string where a few characters of syntactic significance are disallowed.</p>
    <pre class="grammar"><span id="Name">name</span>  ::= [^,()#x20#x9#xD#xA] /* any char except ",", "(", ")" or wsp */</pre></dd><dt id="DataTypeNumber"><span class="SVG-Term">&lt;number&gt;</span></dt><dd>
    <p>Real numbers are specified in one of two ways.  When used in a style sheet,
    a &lt;number&gt; is defined as follows:</p>
    <pre class="grammar"><span id="Number">number</span> ::= <a href="#Integer">integer</a>
           | [+-]? [0-9]* "." [0-9]+</pre>
    <p>This syntax is the same as the definition in CSS
      ([<a href="refs.html#ref-CSS2">CSS2</a>], section 4.3.1).</p>
    <p>When used in an SVG attribute,
    a &lt;number&gt; is defined differently, to allow numbers with large magnitudes
    to be specified more concisely:</p>
    <pre class="grammar">number ::= <a href="#Integer">integer</a> ([Ee] <a href="#DataTypeInteger">integer</a>)?
           | [+-]? [0-9]* "." [0-9]+ ([Ee] <a href="#DataTypeInteger">integer</a>)?</pre>
    <p>Within the SVG DOM, a &lt;number&gt; is represented as a
      <span class="DOMInterfaceName">float</span>, <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> or a
      <a class="idlinterface" href="types.html#InterfaceSVGAnimatedNumber">SVGAnimatedNumber</a>.</p>
  </dd><dt id="DataTypeNumberOptionalNumber"><span class="SVG-Term">&lt;number-optional-number&gt;</span></dt><dd>
    <p>A pair of <a href="types.html#DataTypeNumber">&lt;number&gt;</a>s, where the second
    <a href="types.html#DataTypeNumber">&lt;number&gt;</a> is optional.</p>
    <pre class="grammar">
number-optional-number ::= <a href="#DataTypeNumber">number</a>
                           | <a href="#DataTypeNumber">number</a> <a href="#CommaWSP">comma-wsp</a> <a href="#DataTypeNumber">number</a></pre>
    <p>In the SVG DOM, a &lt;number-optional-number&gt; is represented
    using a pair of <a class="idlinterface" href="types.html#InterfaceSVGAnimatedInteger">SVGAnimatedInteger</a> or <a class="idlinterface" href="types.html#InterfaceSVGAnimatedNumber">SVGAnimatedNumber</a>
    objects.</p>
  </dd><dt id="DataTypePaint"><span class="SVG-Term">&lt;paint&gt;</span></dt><dd>
    <p>The values for properties <a href="painting.html#FillProperty"><span class="prop-name">‘fill’</span></a> and <a href="painting.html#StrokeProperty"><span class="prop-name">‘stroke’</span></a>
    are specifications of the type of paint to use when filling or stroking
    a given graphics element. The available options and syntax for
    &lt;paint&gt; are described in
    <a href="painting.html#SpecifyingPaint">Specifying paint</a>.</p>
    <p>Within the SVG DOM, &lt;paint&gt; values are represented using
    <a class="idlinterface" href="painting.html#InterfaceSVGPaint">SVGPaint</a> objects.</p>
  </dd><dt id="DataTypePercentage"><span class="SVG-Term">&lt;percentage&gt;</span></dt><dd>
    <p>Percentages are specified as a number followed by a "%" character:</p>
    <pre class="grammar"><span id="Percentage">percentage</span> ::= <a href="#Number">number</a> "%"</pre>
    <p>Note that the definition of <a href="#Number">&lt;number&gt;</a> depends
    on whether the percentage is specified in a style sheet or in an
    attribute that is not also a <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a>.</p>
    <p>Percentage values are always relative to another value, for example a
    length. Each attribute or <a href="intro.html#TermProperty"><span class="svg-term">property</span></a> that allows percentages also
    defines the reference distance measurement to which the percentage
    refers.</p>
    <p>Within the SVG DOM, a &lt;percentage&gt; is
    represented using an <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> or <a class="idlinterface" href="types.html#InterfaceSVGAnimatedNumber">SVGAnimatedNumber</a>
    object.</p>
  </dd><dt id="DataTypeTime"><span class="SVG-Term">&lt;time&gt;</span></dt><dd>
    <p>A time value is a &lt;number&gt; immediately followed by a time unit
      identifier. The time unit identifiers are: </p>
    <ul><li><span class="attr-value">ms</span>: milliseconds</li><li><span class="attr-value">s</span>: seconds</li></ul>
    <p>In the SVG DOM, &lt;time&gt; values are represented
      using the <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSPrimitiveValue">CSSPrimitiveValue</a>
      interface defined in <a href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html">Document Object Model CSS</a>
      ([<a href="refs.html#ref-DOM2STYLE">DOM2STYLE</a>], section 2.2).</p>
  </dd><dt id="DataTypeTransformList"><span class="SVG-Term">&lt;transform-list&gt;</span></dt><dd>
    <p>A &lt;transform-list&gt; is used to specify a list of coordinate system
    transformations.  A detailed description of the possible values for a
    &lt;transform-list&gt; is given in
    <a href="coords.html#TransformAttribute">Modifying the User Coordinate
      System: the transform attribute</a>.</p>
    <p>Within the SVG DOM, a &lt;transform-list&gt; value is represented using
    an <a class="idlinterface" href="coords.html#InterfaceSVGTransformList">SVGTransformList</a> or <a class="idlinterface" href="coords.html#InterfaceSVGAnimatedTransformList">SVGAnimatedTransformList</a>
    object.</p>
  </dd><dt id="DataTypeXML-Name"><span class="SVG-Term">&lt;XML-Name&gt;</span></dt><dd>
    <p>
      An XML name, as defined by the
      <a href="http://www.w3.org/TR/2006/REC-xml-20060816/#NT-Name">Name</a>
      production in
      <a href="http://www.w3.org/TR/2006/REC-xml-20060816/"><cite>Extensible Markup Language (XML) 1.0</cite></a> ([<a href="refs.html#ref-XML10">XML10</a>], section 2.3).
    </p>
  </dd></dl>

<h2 id="Precision">4.3 Real number precision</h2>

<p>Unless stated otherwise for a particular attribute or <a href="intro.html#TermProperty"><span class="svg-term">property</span></a>, a
<a href="types.html#DataTypeNumber">&lt;number&gt;</a> has the capacity
for at least a single-precision floating point number and has a range (at a
minimum) of -3.4e+38F to +3.4e+38F.</p>

<p>It is recommended that higher precision floating point
storage and computation be performed on operations such as
coordinate system transformations to provide the best
possible precision and to prevent round-off errors.</p>

<p><a href="conform.html#ConformingHighQualitySVGViewers">Conforming
High-Quality SVG Viewers</a> are required to use at least double-precision
floating point for intermediate calculations on certain numerical operations.</p>

    <h2 id="ColorKeywords">4.4 Recognized color keyword names</h2>
    <p>The following is the list of recognized color keywords that
    can be used as a keyword value for data type <a href="types.html#DataTypeColor">&lt;color&gt;</a>:</p>
    <table class="color-keywords" summary="color keywords"><tr><td>
          <table summary="color keywords part 1"><tr><td><div class="colorpatch aliceblue"/></td><td><span class="prop-value">aliceblue</span></td><td><span class="color-keyword-value">rgb(240, 248, 255)</span></td></tr><tr><td><div class="colorpatch antiquewhite"/></td><td><span class="prop-value">antiquewhite</span></td><td><span class="color-keyword-value">rgb(250, 235, 215)</span></td></tr><tr><td><div class="colorpatch aqua"/></td><td><span class="prop-value">aqua</span></td><td><span class="color-keyword-value">rgb( 0, 255, 255)</span></td></tr><tr><td><div class="colorpatch aquamarine"/></td><td><span class="prop-value">aquamarine</span></td><td><span class="color-keyword-value">rgb(127, 255, 212)</span></td></tr><tr><td><div class="colorpatch azure"/></td><td><span class="prop-value">azure</span></td><td><span class="color-keyword-value">rgb(240, 255, 255)</span></td></tr><tr><td><div class="colorpatch beige"/></td><td><span class="prop-value">beige</span></td><td><span class="color-keyword-value">rgb(245, 245, 220)</span></td></tr><tr><td><div class="colorpatch bisque"/></td><td><span class="prop-value">bisque</span></td><td><span class="color-keyword-value">rgb(255, 228, 196)</span></td></tr><tr><td><div class="colorpatch black"/></td><td><span class="prop-value">black</span></td><td><span class="color-keyword-value">rgb( 0, 0, 0)</span></td></tr><tr><td><div class="colorpatch blanchedalmond"/></td><td><span class="prop-value">blanchedalmond</span></td><td><span class="color-keyword-value">rgb(255, 235, 205)</span></td></tr><tr><td><div class="colorpatch blue"/></td><td><span class="prop-value">blue</span></td><td><span class="color-keyword-value">rgb( 0, 0, 255)</span></td></tr><tr><td><div class="colorpatch blueviolet"/></td><td><span class="prop-value">blueviolet</span></td><td><span class="color-keyword-value">rgb(138, 43, 226)</span></td></tr><tr><td><div class="colorpatch brown"/></td><td><span class="prop-value">brown</span></td><td><span class="color-keyword-value">rgb(165, 42, 42)</span></td></tr><tr><td><div class="colorpatch burlywood"/></td><td><span class="prop-value">burlywood</span></td><td><span class="color-keyword-value">rgb(222, 184, 135)</span></td></tr><tr><td><div class="colorpatch cadetblue"/></td><td><span class="prop-value">cadetblue</span></td><td><span class="color-keyword-value">rgb( 95, 158, 160)</span></td></tr><tr><td><div class="colorpatch chartreuse"/></td><td><span class="prop-value">chartreuse</span></td><td><span class="color-keyword-value">rgb(127, 255, 0)</span></td></tr><tr><td><div class="colorpatch chocolate"/></td><td><span class="prop-value">chocolate</span></td><td><span class="color-keyword-value">rgb(210, 105, 30)</span></td></tr><tr><td><div class="colorpatch coral"/></td><td><span class="prop-value">coral</span></td><td><span class="color-keyword-value">rgb(255, 127, 80)</span></td></tr><tr><td><div class="colorpatch cornflowerblue"/></td><td><span class="prop-value">cornflowerblue</span></td><td><span class="color-keyword-value">rgb(100, 149, 237)</span></td></tr><tr><td><div class="colorpatch cornsilk"/></td><td><span class="prop-value">cornsilk</span></td><td><span class="color-keyword-value">rgb(255, 248, 220)</span></td></tr><tr><td><div class="colorpatch crimson"/></td><td><span class="prop-value">crimson</span></td><td><span class="color-keyword-value">rgb(220, 20, 60)</span></td></tr><tr><td><div class="colorpatch cyan"/></td><td><span class="prop-value">cyan</span></td><td><span class="color-keyword-value">rgb( 0, 255, 255)</span></td></tr><tr><td><div class="colorpatch darkblue"/></td><td><span class="prop-value">darkblue</span></td><td><span class="color-keyword-value">rgb( 0, 0, 139)</span></td></tr><tr><td><div class="colorpatch darkcyan"/></td><td><span class="prop-value">darkcyan</span></td><td><span class="color-keyword-value">rgb( 0, 139, 139)</span></td></tr><tr><td><div class="colorpatch darkgoldenrod"/></td><td><span class="prop-value">darkgoldenrod</span></td><td><span class="color-keyword-value">rgb(184, 134, 11)</span></td></tr><tr><td><div class="colorpatch darkgray"/></td><td><span class="prop-value">darkgray</span></td><td><span class="color-keyword-value">rgb(169, 169, 169)</span></td></tr><tr><td><div class="colorpatch darkgreen"/></td><td><span class="prop-value">darkgreen</span></td><td><span class="color-keyword-value">rgb( 0, 100, 0)</span></td></tr><tr><td><div class="colorpatch darkgrey"/></td><td><span class="prop-value">darkgrey</span></td><td><span class="color-keyword-value">rgb(169, 169, 169)</span></td></tr><tr><td><div class="colorpatch darkkhaki"/></td><td><span class="prop-value">darkkhaki</span></td><td><span class="color-keyword-value">rgb(189, 183, 107)</span></td></tr><tr><td><div class="colorpatch darkmagenta"/></td><td><span class="prop-value">darkmagenta</span></td><td><span class="color-keyword-value">rgb(139, 0, 139)</span></td></tr><tr><td><div class="colorpatch darkolivegreen"/></td><td><span class="prop-value">darkolivegreen</span></td><td><span class="color-keyword-value">rgb( 85, 107, 47)</span></td></tr><tr><td><div class="colorpatch darkorange"/></td><td><span class="prop-value">darkorange</span></td><td><span class="color-keyword-value">rgb(255, 140, 0)</span></td></tr><tr><td><div class="colorpatch darkorchid"/></td><td><span class="prop-value">darkorchid</span></td><td><span class="color-keyword-value">rgb(153, 50, 204)</span></td></tr><tr><td><div class="colorpatch darkred"/></td><td><span class="prop-value">darkred</span></td><td><span class="color-keyword-value">rgb(139, 0, 0)</span></td></tr><tr><td><div class="colorpatch darksalmon"/></td><td><span class="prop-value">darksalmon</span></td><td><span class="color-keyword-value">rgb(233, 150, 122)</span></td></tr><tr><td><div class="colorpatch darkseagreen"/></td><td><span class="prop-value">darkseagreen</span></td><td><span class="color-keyword-value">rgb(143, 188, 143)</span></td></tr><tr><td><div class="colorpatch darkslateblue"/></td><td><span class="prop-value">darkslateblue</span></td><td><span class="color-keyword-value">rgb( 72, 61, 139)</span></td></tr><tr><td><div class="colorpatch darkslategray"/></td><td><span class="prop-value">darkslategray</span></td><td><span class="color-keyword-value">rgb( 47, 79, 79)</span></td></tr><tr><td><div class="colorpatch darkslategrey"/></td><td><span class="prop-value">darkslategrey</span></td><td><span class="color-keyword-value">rgb( 47, 79, 79)</span></td></tr><tr><td><div class="colorpatch darkturquoise"/></td><td><span class="prop-value">darkturquoise</span></td><td><span class="color-keyword-value">rgb( 0, 206, 209)</span></td></tr><tr><td><div class="colorpatch darkviolet"/></td><td><span class="prop-value">darkviolet</span></td><td><span class="color-keyword-value">rgb(148, 0, 211)</span></td></tr><tr><td><div class="colorpatch deeppink"/></td><td><span class="prop-value">deeppink</span></td><td><span class="color-keyword-value">rgb(255, 20, 147)</span></td></tr><tr><td><div class="colorpatch deepskyblue"/></td><td><span class="prop-value">deepskyblue</span></td><td><span class="color-keyword-value">rgb( 0, 191, 255)</span></td></tr><tr><td><div class="colorpatch dimgray"/></td><td><span class="prop-value">dimgray</span></td><td><span class="color-keyword-value">rgb(105, 105, 105)</span></td></tr><tr><td><div class="colorpatch dimgrey"/></td><td><span class="prop-value">dimgrey</span></td><td><span class="color-keyword-value">rgb(105, 105, 105)</span></td></tr><tr><td><div class="colorpatch dodgerblue"/></td><td><span class="prop-value">dodgerblue</span></td><td><span class="color-keyword-value">rgb( 30, 144, 255)</span></td></tr><tr><td><div class="colorpatch firebrick"/></td><td><span class="prop-value">firebrick</span></td><td><span class="color-keyword-value">rgb(178, 34, 34)</span></td></tr><tr><td><div class="colorpatch floralwhite"/></td><td><span class="prop-value">floralwhite</span></td><td><span class="color-keyword-value">rgb(255, 250, 240)</span></td></tr><tr><td><div class="colorpatch forestgreen"/></td><td><span class="prop-value">forestgreen</span></td><td><span class="color-keyword-value">rgb( 34, 139, 34)</span></td></tr><tr><td><div class="colorpatch fuchsia"/></td><td><span class="prop-value">fuchsia</span></td><td><span class="color-keyword-value">rgb(255, 0, 255)</span></td></tr><tr><td><div class="colorpatch gainsboro"/></td><td><span class="prop-value">gainsboro</span></td><td><span class="color-keyword-value">rgb(220, 220, 220)</span></td></tr><tr><td><div class="colorpatch ghostwhite"/></td><td><span class="prop-value">ghostwhite</span></td><td><span class="color-keyword-value">rgb(248, 248, 255)</span></td></tr><tr><td><div class="colorpatch gold"/></td><td><span class="prop-value">gold</span></td><td><span class="color-keyword-value">rgb(255, 215, 0)</span></td></tr><tr><td><div class="colorpatch goldenrod"/></td><td><span class="prop-value">goldenrod</span></td><td><span class="color-keyword-value">rgb(218, 165, 32)</span></td></tr><tr><td><div class="colorpatch gray"/></td><td><span class="prop-value">gray</span></td><td><span class="color-keyword-value">rgb(128, 128, 128)</span></td></tr><tr><td><div class="colorpatch grey"/></td><td><span class="prop-value">grey</span></td><td><span class="color-keyword-value">rgb(128, 128, 128)</span></td></tr><tr><td><div class="colorpatch green"/></td><td><span class="prop-value">green</span></td><td><span class="color-keyword-value">rgb( 0, 128, 0)</span></td></tr><tr><td><div class="colorpatch greenyellow"/></td><td><span class="prop-value">greenyellow</span></td><td><span class="color-keyword-value">rgb(173, 255, 47)</span></td></tr><tr><td><div class="colorpatch honeydew"/></td><td><span class="prop-value">honeydew</span></td><td><span class="color-keyword-value">rgb(240, 255, 240)</span></td></tr><tr><td><div class="colorpatch hotpink"/></td><td><span class="prop-value">hotpink</span></td><td><span class="color-keyword-value">rgb(255, 105, 180)</span></td></tr><tr><td><div class="colorpatch indianred"/></td><td><span class="prop-value">indianred</span></td><td><span class="color-keyword-value">rgb(205, 92, 92)</span></td></tr><tr><td><div class="colorpatch indigo"/></td><td><span class="prop-value">indigo</span></td><td><span class="color-keyword-value">rgb( 75, 0, 130)</span></td></tr><tr><td><div class="colorpatch ivory"/></td><td><span class="prop-value">ivory</span></td><td><span class="color-keyword-value">rgb(255, 255, 240)</span></td></tr><tr><td><div class="colorpatch khaki"/></td><td><span class="prop-value">khaki</span></td><td><span class="color-keyword-value">rgb(240, 230, 140)</span></td></tr><tr><td><div class="colorpatch lavender"/></td><td><span class="prop-value">lavender</span></td><td><span class="color-keyword-value">rgb(230, 230, 250)</span></td></tr><tr><td><div class="colorpatch lavenderblush"/></td><td><span class="prop-value">lavenderblush</span></td><td><span class="color-keyword-value">rgb(255, 240, 245)</span></td></tr><tr><td><div class="colorpatch lawngreen"/></td><td><span class="prop-value">lawngreen</span></td><td><span class="color-keyword-value">rgb(124, 252, 0)</span></td></tr><tr><td><div class="colorpatch lemonchiffon"/></td><td><span class="prop-value">lemonchiffon</span></td><td><span class="color-keyword-value">rgb(255, 250, 205)</span></td></tr><tr><td><div class="colorpatch lightblue"/></td><td><span class="prop-value">lightblue</span></td><td><span class="color-keyword-value">rgb(173, 216, 230)</span></td></tr><tr><td><div class="colorpatch lightcoral"/></td><td><span class="prop-value">lightcoral</span></td><td><span class="color-keyword-value">rgb(240, 128, 128)</span></td></tr><tr><td><div class="colorpatch lightcyan"/></td><td><span class="prop-value">lightcyan</span></td><td><span class="color-keyword-value">rgb(224, 255, 255)</span></td></tr><tr><td><div class="colorpatch lightgoldenrodyellow"/></td><td><span class="prop-value">lightgoldenrodyellow</span></td><td><span class="color-keyword-value">rgb(250, 250, 210)</span></td></tr><tr><td><div class="colorpatch lightgray"/></td><td><span class="prop-value">lightgray</span></td><td><span class="color-keyword-value">rgb(211, 211, 211)</span></td></tr><tr><td><div class="colorpatch lightgreen"/></td><td><span class="prop-value">lightgreen</span></td><td><span class="color-keyword-value">rgb(144, 238, 144)</span></td></tr><tr><td><div class="colorpatch lightgrey"/></td><td><span class="prop-value">lightgrey</span></td><td><span class="color-keyword-value">rgb(211, 211, 211)</span></td></tr></table>
        </td><td>    </td><td>
          <table summary="color keywords part 2"><tr><td><div class="colorpatch lightpink"/></td><td><span class="prop-value">lightpink</span></td><td><span class="color-keyword-value">rgb(255, 182, 193)</span></td></tr><tr><td><div class="colorpatch lightsalmon"/></td><td><span class="prop-value">lightsalmon</span></td><td><span class="color-keyword-value">rgb(255, 160, 122)</span></td></tr><tr><td><div class="colorpatch lightseagreen"/></td><td><span class="prop-value">lightseagreen</span></td><td><span class="color-keyword-value">rgb( 32, 178, 170)</span></td></tr><tr><td><div class="colorpatch lightskyblue"/></td><td><span class="prop-value">lightskyblue</span></td><td><span class="color-keyword-value">rgb(135, 206, 250)</span></td></tr><tr><td><div class="colorpatch lightslategray"/></td><td><span class="prop-value">lightslategray</span></td><td><span class="color-keyword-value">rgb(119, 136, 153)</span></td></tr><tr><td><div class="colorpatch lightslategrey"/></td><td><span class="prop-value">lightslategrey</span></td><td><span class="color-keyword-value">rgb(119, 136, 153)</span></td></tr><tr><td><div class="colorpatch lightsteelblue"/></td><td><span class="prop-value">lightsteelblue</span></td><td><span class="color-keyword-value">rgb(176, 196, 222)</span></td></tr><tr><td><div class="colorpatch lightyellow"/></td><td><span class="prop-value">lightyellow</span></td><td><span class="color-keyword-value">rgb(255, 255, 224)</span></td></tr><tr><td><div class="colorpatch lime"/></td><td><span class="prop-value">lime</span></td><td><span class="color-keyword-value">rgb( 0, 255, 0)</span></td></tr><tr><td><div class="colorpatch limegreen"/></td><td><span class="prop-value">limegreen</span></td><td><span class="color-keyword-value">rgb( 50, 205, 50)</span></td></tr><tr><td><div class="colorpatch linen"/></td><td><span class="prop-value">linen</span></td><td><span class="color-keyword-value">rgb(250, 240, 230)</span></td></tr><tr><td><div class="colorpatch magenta"/></td><td><span class="prop-value">magenta</span></td><td><span class="color-keyword-value">rgb(255, 0, 255)</span></td></tr><tr><td><div class="colorpatch maroon"/></td><td><span class="prop-value">maroon</span></td><td><span class="color-keyword-value">rgb(128, 0, 0)</span></td></tr><tr><td><div class="colorpatch mediumaquamarine"/></td><td><span class="prop-value">mediumaquamarine</span></td><td><span class="color-keyword-value">rgb(102, 205, 170)</span></td></tr><tr><td><div class="colorpatch mediumblue"/></td><td><span class="prop-value">mediumblue</span></td><td><span class="color-keyword-value">rgb( 0, 0, 205)</span></td></tr><tr><td><div class="colorpatch mediumorchid"/></td><td><span class="prop-value">mediumorchid</span></td><td><span class="color-keyword-value">rgb(186, 85, 211)</span></td></tr><tr><td><div class="colorpatch mediumpurple"/></td><td><span class="prop-value">mediumpurple</span></td><td><span class="color-keyword-value">rgb(147, 112, 219)</span></td></tr><tr><td><div class="colorpatch mediumseagreen"/></td><td><span class="prop-value">mediumseagreen</span></td><td><span class="color-keyword-value">rgb( 60, 179, 113)</span></td></tr><tr><td><div class="colorpatch mediumslateblue"/></td><td><span class="prop-value">mediumslateblue</span></td><td><span class="color-keyword-value">rgb(123, 104, 238)</span></td></tr><tr><td><div class="colorpatch mediumspringgreen"/></td><td><span class="prop-value">mediumspringgreen</span></td><td><span class="color-keyword-value">rgb( 0, 250, 154)</span></td></tr><tr><td><div class="colorpatch mediumturquoise"/></td><td><span class="prop-value">mediumturquoise</span></td><td><span class="color-keyword-value">rgb( 72, 209, 204)</span></td></tr><tr><td><div class="colorpatch mediumvioletred"/></td><td><span class="prop-value">mediumvioletred</span></td><td><span class="color-keyword-value">rgb(199, 21, 133)</span></td></tr><tr><td><div class="colorpatch midnightblue"/></td><td><span class="prop-value">midnightblue</span></td><td><span class="color-keyword-value">rgb( 25, 25, 112)</span></td></tr><tr><td><div class="colorpatch mintcream"/></td><td><span class="prop-value">mintcream</span></td><td><span class="color-keyword-value">rgb(245, 255, 250)</span></td></tr><tr><td><div class="colorpatch mistyrose"/></td><td><span class="prop-value">mistyrose</span></td><td><span class="color-keyword-value">rgb(255, 228, 225)</span></td></tr><tr><td><div class="colorpatch moccasin"/></td><td><span class="prop-value">moccasin</span></td><td><span class="color-keyword-value">rgb(255, 228, 181)</span></td></tr><tr><td><div class="colorpatch navajowhite"/></td><td><span class="prop-value">navajowhite</span></td><td><span class="color-keyword-value">rgb(255, 222, 173)</span></td></tr><tr><td><div class="colorpatch navy"/></td><td><span class="prop-value">navy</span></td><td><span class="color-keyword-value">rgb( 0, 0, 128)</span></td></tr><tr><td><div class="colorpatch oldlace"/></td><td><span class="prop-value">oldlace</span></td><td><span class="color-keyword-value">rgb(253, 245, 230)</span></td></tr><tr><td><div class="colorpatch olive"/></td><td><span class="prop-value">olive</span></td><td><span class="color-keyword-value">rgb(128, 128, 0)</span></td></tr><tr><td><div class="colorpatch olivedrab"/></td><td><span class="prop-value">olivedrab</span></td><td><span class="color-keyword-value">rgb(107, 142, 35)</span></td></tr><tr><td><div class="colorpatch orange"/></td><td><span class="prop-value">orange</span></td><td><span class="color-keyword-value">rgb(255, 165, 0)</span></td></tr><tr><td><div class="colorpatch orangered"/></td><td><span class="prop-value">orangered</span></td><td><span class="color-keyword-value">rgb(255, 69, 0)</span></td></tr><tr><td><div class="colorpatch orchid"/></td><td><span class="prop-value">orchid</span></td><td><span class="color-keyword-value">rgb(218, 112, 214)</span></td></tr><tr><td><div class="colorpatch palegoldenrod"/></td><td><span class="prop-value">palegoldenrod</span></td><td><span class="color-keyword-value">rgb(238, 232, 170)</span></td></tr><tr><td><div class="colorpatch palegreen"/></td><td><span class="prop-value">palegreen</span></td><td><span class="color-keyword-value">rgb(152, 251, 152)</span></td></tr><tr><td><div class="colorpatch paleturquoise"/></td><td><span class="prop-value">paleturquoise</span></td><td><span class="color-keyword-value">rgb(175, 238, 238)</span></td></tr><tr><td><div class="colorpatch palevioletred"/></td><td><span class="prop-value">palevioletred</span></td><td><span class="color-keyword-value">rgb(219, 112, 147)</span></td></tr><tr><td><div class="colorpatch papayawhip"/></td><td><span class="prop-value">papayawhip</span></td><td><span class="color-keyword-value">rgb(255, 239, 213)</span></td></tr><tr><td><div class="colorpatch peachpuff"/></td><td><span class="prop-value">peachpuff</span></td><td><span class="color-keyword-value">rgb(255, 218, 185)</span></td></tr><tr><td><div class="colorpatch peru"/></td><td><span class="prop-value">peru</span></td><td><span class="color-keyword-value">rgb(205, 133, 63)</span></td></tr><tr><td><div class="colorpatch pink"/></td><td><span class="prop-value">pink</span></td><td><span class="color-keyword-value">rgb(255, 192, 203)</span></td></tr><tr><td><div class="colorpatch plum"/></td><td><span class="prop-value">plum</span></td><td><span class="color-keyword-value">rgb(221, 160, 221)</span></td></tr><tr><td><div class="colorpatch powderblue"/></td><td><span class="prop-value">powderblue</span></td><td><span class="color-keyword-value">rgb(176, 224, 230)</span></td></tr><tr><td><div class="colorpatch purple"/></td><td><span class="prop-value">purple</span></td><td><span class="color-keyword-value">rgb(128, 0, 128)</span></td></tr><tr><td><div class="colorpatch red"/></td><td><span class="prop-value">red</span></td><td><span class="color-keyword-value">rgb(255, 0, 0)</span></td></tr><tr><td><div class="colorpatch rosybrown"/></td><td><span class="prop-value">rosybrown</span></td><td><span class="color-keyword-value">rgb(188, 143, 143)</span></td></tr><tr><td><div class="colorpatch royalblue"/></td><td><span class="prop-value">royalblue</span></td><td><span class="color-keyword-value">rgb( 65, 105, 225)</span></td></tr><tr><td><div class="colorpatch saddlebrown"/></td><td><span class="prop-value">saddlebrown</span></td><td><span class="color-keyword-value">rgb(139, 69, 19)</span></td></tr><tr><td><div class="colorpatch salmon"/></td><td><span class="prop-value">salmon</span></td><td><span class="color-keyword-value">rgb(250, 128, 114)</span></td></tr><tr><td><div class="colorpatch sandybrown"/></td><td><span class="prop-value">sandybrown</span></td><td><span class="color-keyword-value">rgb(244, 164, 96)</span></td></tr><tr><td><div class="colorpatch seagreen"/></td><td><span class="prop-value">seagreen</span></td><td><span class="color-keyword-value">rgb( 46, 139, 87)</span></td></tr><tr><td><div class="colorpatch seashell"/></td><td><span class="prop-value">seashell</span></td><td><span class="color-keyword-value">rgb(255, 245, 238)</span></td></tr><tr><td><div class="colorpatch sienna"/></td><td><span class="prop-value">sienna</span></td><td><span class="color-keyword-value">rgb(160, 82, 45)</span></td></tr><tr><td><div class="colorpatch silver"/></td><td><span class="prop-value">silver</span></td><td><span class="color-keyword-value">rgb(192, 192, 192)</span></td></tr><tr><td><div class="colorpatch skyblue"/></td><td><span class="prop-value">skyblue</span></td><td><span class="color-keyword-value">rgb(135, 206, 235)</span></td></tr><tr><td><div class="colorpatch slateblue"/></td><td><span class="prop-value">slateblue</span></td><td><span class="color-keyword-value">rgb(106, 90, 205)</span></td></tr><tr><td><div class="colorpatch slategray"/></td><td><span class="prop-value">slategray</span></td><td><span class="color-keyword-value">rgb(112, 128, 144)</span></td></tr><tr><td><div class="colorpatch slategrey"/></td><td><span class="prop-value">slategrey</span></td><td><span class="color-keyword-value">rgb(112, 128, 144)</span></td></tr><tr><td><div class="colorpatch snow"/></td><td><span class="prop-value">snow</span></td><td><span class="color-keyword-value">rgb(255, 250, 250)</span></td></tr><tr><td><div class="colorpatch springgreen"/></td><td><span class="prop-value">springgreen</span></td><td><span class="color-keyword-value">rgb( 0, 255, 127)</span></td></tr><tr><td><div class="colorpatch steelblue"/></td><td><span class="prop-value">steelblue</span></td><td><span class="color-keyword-value">rgb( 70, 130, 180)</span></td></tr><tr><td><div class="colorpatch tan"/>
              </td><td><span class="prop-value">tan</span></td><td><span class="color-keyword-value">rgb(210, 180, 140)</span></td></tr><tr><td><div class="colorpatch teal"/></td><td><span class="prop-value">teal</span></td><td><span class="color-keyword-value">rgb( 0, 128, 128)</span></td></tr><tr><td><div class="colorpatch thistle"/></td><td><span class="prop-value">thistle</span></td><td><span class="color-keyword-value">rgb(216, 191, 216)</span></td></tr><tr><td><div class="colorpatch tomato"/></td><td><span class="prop-value">tomato</span></td><td><span class="color-keyword-value">rgb(255, 99, 71)</span></td></tr><tr><td><div class="colorpatch turquoise"/></td><td><span class="prop-value">turquoise</span></td><td><span class="color-keyword-value">rgb( 64, 224, 208)</span></td></tr><tr><td><div class="colorpatch violet"/></td><td><span class="prop-value">violet</span></td><td><span class="color-keyword-value">rgb(238, 130, 238)</span></td></tr><tr><td><div class="colorpatch wheat"/></td><td><span class="prop-value">wheat</span></td><td><span class="color-keyword-value">rgb(245, 222, 179)</span></td></tr><tr><td><div class="colorpatch white"/></td><td><span class="prop-value">white</span></td><td><span class="color-keyword-value">rgb(255, 255, 255)</span></td></tr><tr><td><div class="colorpatch whitesmoke"/>
              </td><td><span class="prop-value">whitesmoke</span></td><td><span class="color-keyword-value">rgb(245, 245, 245)</span></td></tr><tr><td><div class="colorpatch yellow"/>
              </td><td><span class="prop-value">yellow</span></td><td><span class="color-keyword-value">rgb(255, 255, 0)</span></td></tr><tr><td><div class="colorpatch yellowgreen"/>
              </td><td><span class="prop-value">yellowgreen</span></td><td><span class="color-keyword-value">rgb(154, 205, 50)</span></td></tr><tr><td/><td><span class="prop-value"> </span></td><td> </td></tr></table>
        </td></tr></table>

<h2 id="BasicDOMInterfaces">4.5 Basic DOM interfaces</h2>

<h3 id="InterfaceSVGElement">4.5.1 Interface SVGElement</h3>


 All of the SVG DOM interfaces that correspond directly to elements in the
 SVG language (such as the <a class="idlinterface" href="paths.html#InterfaceSVGPathElement">SVGPathElement</a> interface for the
 <a href="paths.html#PathElement"><span class="element-name">‘path’</span></a> element) derive from the <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> interface.
<pre class="idl">interface <b>SVGElement</b> : <a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-745549614">Element</a> {
           attribute DOMString <a href="types.html#__svg__SVGElement__id">id</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
           attribute DOMString <a href="types.html#__svg__SVGElement__xmlbase">xmlbase</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute <a class="idlinterface" href="struct.html#InterfaceSVGSVGElement">SVGSVGElement</a> <a href="types.html#__svg__SVGElement__ownerSVGElement">ownerSVGElement</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> <a href="types.html#__svg__SVGElement__viewportElement">viewportElement</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGElement__id" class="attribute first-child"><b>id</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 The value of the <a href="struct.html#IDAttribute"><span class="attr-name">‘id’</span></a> attribute on the given element, or the
 empty string if <a href="struct.html#IDAttribute"><span class="attr-name">‘id’</span></a> is not present.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an attempt
   to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGElement__xmlbase" class="attribute"><b>xmlbase</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#XMLBaseAttribute"><span class="attr-name">‘xml:base’</span></a> on the given element.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an attempt
   to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGElement__ownerSVGElement" class="attribute"><b>ownerSVGElement</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="struct.html#InterfaceSVGSVGElement">SVGSVGElement</a>)</span></dt><dd class="attribute"><div>
 The nearest ancestor <a href="struct.html#SVGElement"><span class="element-name">‘svg’</span></a> element. Null if the given element is
 the <a href="intro.html#TermOutermostSVGElement"><span class="svg-term">outermost svg element</span></a>.
</div></dd>
<dt id="__svg__SVGElement__viewportElement" class="attribute"><b>viewportElement</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a>)</span></dt><dd class="attribute"><div>
 The element which established the current viewport. Often, the nearest
 ancestor <a href="struct.html#SVGElement"><span class="element-name">‘svg’</span></a> element. Null if the given element is the
 <a href="intro.html#TermOutermostSVGElement"><span class="svg-term">outermost svg element</span></a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedBoolean">4.5.2 Interface SVGAnimatedBoolean</h3>


 Used for attributes of type boolean which can be animated.
<pre class="idl">interface <b>SVGAnimatedBoolean</b> {
           attribute boolean <a href="types.html#__svg__SVGAnimatedBoolean__baseVal">baseVal</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute boolean <a href="types.html#__svg__SVGAnimatedBoolean__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedBoolean__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (boolean)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAnimatedBoolean__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly boolean)</span></dt><dd class="attribute"><div>
 If the given attribute or property is being animated, contains the
 current animated value of the attribute or property. If the given
 attribute or property is not currently being animated, contains the
 same value as <a class="idlattr" href="#__svg__SVGAnimatedBoolean__baseVal">baseVal</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedString">4.5.3 Interface SVGAnimatedString</h3>


 Used for attributes of type DOMString which can be animated.
<pre class="idl">interface <b>SVGAnimatedString</b> {
           attribute DOMString <a href="types.html#__svg__SVGAnimatedString__baseVal">baseVal</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute DOMString <a href="types.html#__svg__SVGAnimatedString__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedString__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAnimatedString__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly DOMString)</span></dt><dd class="attribute"><div>
 If the given attribute or property is being animated, contains the
 current animated value of the attribute or property. If the given
 attribute or property is not currently being animated, contains the
 same value as <a class="idlattr" href="#__svg__SVGAnimatedString__baseVal">baseVal</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGStringList">4.5.4 Interface SVGStringList</h3>


 <p>This interface defines a list of DOMString values.</p>

 <p><a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a> has the same attributes and methods as other
 SVGxxxList interfaces. Implementers may consider using a single base class
 to implement the various SVGxxxList interfaces.</p>
<pre class="idl">interface <b>SVGStringList</b> {

  readonly attribute unsigned long <a href="types.html#__svg__SVGStringList__numberOfItems">numberOfItems</a>;

  void <a href="types.html#__svg__SVGStringList__clear">clear</a>() raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  DOMString <a href="types.html#__svg__SVGStringList__initialize">initialize</a>(in DOMString newItem) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  DOMString <a href="types.html#__svg__SVGStringList__getItem">getItem</a>(in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  DOMString <a href="types.html#__svg__SVGStringList__insertItemBefore">insertItemBefore</a>(in DOMString newItem, in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  DOMString <a href="types.html#__svg__SVGStringList__replaceItem">replaceItem</a>(in DOMString newItem, in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  DOMString <a href="types.html#__svg__SVGStringList__removeItem">removeItem</a>(in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  DOMString <a href="types.html#__svg__SVGStringList__appendItem">appendItem</a>(in DOMString newItem) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGStringList__numberOfItems" class="attribute first-child"><b>numberOfItems</b><span class="idl-type-parenthetical"> (readonly unsigned long)</span></dt><dd class="attribute"><div>
 The number of items in the list.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGStringList__clear" class="operation first-child">void <b>clear</b>()</dt><dd class="operation"><div>
 Clears all existing current items from the list, with the result being
 an empty list.

</div><dl class="operation"><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGStringList__initialize" class="operation">DOMString <b>initialize</b>(in DOMString <var>newItem</var>)</dt><dd class="operation"><div>
 Clears all existing current items from the list and re-initializes the
 list to hold the single item specified by the parameter.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>newItem</var></div> <div> The item which should become the only member of the list.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The item being inserted into the list.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGStringList__getItem" class="operation">DOMString <b>getItem</b>(in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Returns the specified item from the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned long <var>index</var></div> <div> The index of the item from the list which is to be
   returned.  The first item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The selected item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGStringList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGStringList__insertItemBefore" class="operation">DOMString <b>insertItemBefore</b>(in DOMString <var>newItem</var>, in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Inserts a new item into the list at the specified position. The first
 item is number 0.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>newItem</var></div> <div> The item which is to be inserted into the list.
</div></li><li class="parameter"><div>unsigned long <var>index</var></div> <div> The index of the item before which the new item is to be
   inserted. The first item is number 0.  If the index is equal to 0,
   then the new item is inserted at the front of the list. If the index
   is greater than or equal to <a class="idlattr" href="#__svg__SVGStringList__numberOfItems">numberOfItems</a>, then the new item is
   appended to the end of the list.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGStringList__replaceItem" class="operation">DOMString <b>replaceItem</b>(in DOMString <var>newItem</var>, in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Replaces an existing item in the list with a new item.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>newItem</var></div> <div> The item which is to be inserted into the list.
</div></li><li class="parameter"><div>unsigned long <var>index</var></div> <div> The index of the item which is to be replaced. The first
   item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGStringList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGStringList__removeItem" class="operation">DOMString <b>removeItem</b>(in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Removes an existing item from the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned long <var>index</var></div> <div> The index of the item which is to be removed. The first
   item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The removed item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGStringList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGStringList__appendItem" class="operation">DOMString <b>appendItem</b>(in DOMString <var>newItem</var>)</dt><dd class="operation"><div>
 Inserts a new item at the end of the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>newItem</var></div> <div> The item which is to be inserted. The first item is
   number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedEnumeration">4.5.5 Interface SVGAnimatedEnumeration</h3>


 Used for attributes whose value must be a constant from a particular
 enumeration and which can be animated.
<pre class="idl">interface <b>SVGAnimatedEnumeration</b> {
           attribute unsigned short <a href="types.html#__svg__SVGAnimatedEnumeration__baseVal">baseVal</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute unsigned short <a href="types.html#__svg__SVGAnimatedEnumeration__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedEnumeration__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAnimatedEnumeration__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly unsigned short)</span></dt><dd class="attribute"><div>
 If the given attribute or property is being animated, contains the
 current animated value of the attribute or property. If the given
 attribute or property is not currently being animated, contains the
 same value as <a class="idlattr" href="#__svg__SVGAnimatedEnumeration__baseVal">baseVal</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedInteger">4.5.6 Interface SVGAnimatedInteger</h3>


 Used for attributes of basic type
 <a href="types.html#DataTypeInteger">&lt;integer&gt;</a> which can be
 animated.
<pre class="idl">interface <b>SVGAnimatedInteger</b> {
           attribute long <a href="types.html#__svg__SVGAnimatedInteger__baseVal">baseVal</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute long <a href="types.html#__svg__SVGAnimatedInteger__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedInteger__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (long)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAnimatedInteger__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly long)</span></dt><dd class="attribute"><div>
 If the given attribute or property is being animated, contains the
 current animated value of the attribute or property. If the given
 attribute or property is not currently being animated, contains the
 same value as <a class="idlattr" href="#__svg__SVGAnimatedInteger__baseVal">baseVal</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGNumber">4.5.7 Interface SVGNumber</h3>


 Used for attributes of basic type
 <a href="types.html#DataTypeNumber">&lt;number&gt;</a>.
<pre class="idl">interface <b>SVGNumber</b> {
  attribute float <a href="types.html#__svg__SVGNumber__value">value</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGNumber__value" class="attribute first-child"><b>value</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The value of the given attribute.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedNumber">4.5.8 Interface SVGAnimatedNumber</h3>


 Used for attributes of basic type
 <a href="types.html#DataTypeNumber">&lt;number&gt;</a> which can be
 animated.
<pre class="idl">interface <b>SVGAnimatedNumber</b> {
           attribute float <a href="types.html#__svg__SVGAnimatedNumber__baseVal">baseVal</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute float <a href="types.html#__svg__SVGAnimatedNumber__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedNumber__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAnimatedNumber__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly float)</span></dt><dd class="attribute"><div>
 If the given attribute or property is being animated, contains the
 current animated value of the attribute or property. If the given
 attribute or property is not currently being animated, contains the
 same value as <a class="idlattr" href="#__svg__SVGAnimatedNumber__baseVal">baseVal</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGNumberList">4.5.9 Interface SVGNumberList</h3>


 <p>This interface defines a list of SVGNumber objects.</p>

 <p><a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> has the same attributes and methods as other
 SVGxxxList interfaces. Implementers may consider using a single base class
 to implement the various SVGxxxList interfaces.</p>

 <p id="ReadOnlyNumberList">An <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> object can be designated as <em>read only</em>,
 which means that attempts to modify the object will result in an exception
 being thrown, as described below.</p>
<pre class="idl">interface <b>SVGNumberList</b> {

  readonly attribute unsigned long <a href="types.html#__svg__SVGNumberList__numberOfItems">numberOfItems</a>;

  void <a href="types.html#__svg__SVGNumberList__clear">clear</a>() raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <a href="types.html#__svg__SVGNumberList__initialize">initialize</a>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> newItem) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <a href="types.html#__svg__SVGNumberList__getItem">getItem</a>(in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <a href="types.html#__svg__SVGNumberList__insertItemBefore">insertItemBefore</a>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> newItem, in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <a href="types.html#__svg__SVGNumberList__replaceItem">replaceItem</a>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> newItem, in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <a href="types.html#__svg__SVGNumberList__removeItem">removeItem</a>(in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <a href="types.html#__svg__SVGNumberList__appendItem">appendItem</a>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> newItem) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGNumberList__numberOfItems" class="attribute first-child"><b>numberOfItems</b><span class="idl-type-parenthetical"> (readonly unsigned long)</span></dt><dd class="attribute"><div>
 The number of items in the list.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGNumberList__clear" class="operation first-child">void <b>clear</b>()</dt><dd class="operation"><div>
 Clears all existing current items from the list, with the result being
 an empty list.

</div><dl class="operation"><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyNumberList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGNumberList__initialize" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <b>initialize</b>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var>)</dt><dd class="operation"><div>
 Clears all existing current items from the list and re-initializes the
 list to hold the single item specified by the parameter.  If the inserted
 item is already in a list, it is removed from its previous list before
 it is inserted into this list.  The inserted item is the item itself and
 not a copy. 

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var></div> <div> The item which should become the only member of the list.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The item being inserted into the list.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyNumberList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGNumberList__getItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <b>getItem</b>(in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Returns the specified item from the list.  The returned item is the
 item itself and not a copy.  Any changes made to the item are
 immediately reflected in the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned long <var>index</var></div> <div> The index of the item from the list which is to be
   returned.  The first item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The selected item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGNumberList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGNumberList__insertItemBefore" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <b>insertItemBefore</b>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var>, in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Inserts a new item into the list at the specified position. The first
 item is number 0. If <var>newItem</var> is already in a list, it is
 removed from its previous list before it is inserted into this list.
 The inserted item is the item itself and not a copy. If the item is
 already in this list, note that the index of the item to insert
 before is <i>before</i> the removal of the item.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var></div> <div> The item which is to be inserted into the list.
</div></li><li class="parameter"><div>unsigned long <var>index</var></div> <div> The index of the item before which the new item is to be
   inserted. The first item is number 0.  If the index is equal to 0,
   then the new item is inserted at the front of the list. If the index
   is greater than or equal to <a class="idlattr" href="#__svg__SVGNumberList__numberOfItems">numberOfItems</a>, then the new item is
   appended to the end of the list.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyNumberList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGNumberList__replaceItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <b>replaceItem</b>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var>, in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Replaces an existing item in the list with a new item. If
 <var>newItem</var> is already in a list, it is removed from its
 previous list before it is inserted into this list.  The inserted item
 is the item itself and not a copy.  If the item is already in this
 list, note that the index of the item to replace is <i>before</i>
 the removal of the item.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var></div> <div> The item which is to be inserted into the list.
</div></li><li class="parameter"><div>unsigned long <var>index</var></div> <div> The index of the item which is to be replaced. The first
   item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyNumberList">read only</a>.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGNumberList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGNumberList__removeItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <b>removeItem</b>(in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Removes an existing item from the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned long <var>index</var></div> <div> The index of the item which is to be removed. The first
   item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The removed item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGNumberList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGNumberList__appendItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <b>appendItem</b>(in <a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var>)</dt><dd class="operation"><div>
 Inserts a new item at the end of the list. If <var>newItem</var> is
 already in a list, it is removed from its previous list before it is
 inserted into this list.  The inserted item is the item itself and
 not a copy.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGNumber">SVGNumber</a> <var>newItem</var></div> <div> The item which is to be inserted. The first item is
   number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   cannot be modified.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedNumberList">4.5.10 Interface SVGAnimatedNumberList</h3>


 Used for attributes which take a list of numbers and which can be animated.
<pre class="idl">interface <b>SVGAnimatedNumberList</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> <a href="types.html#__svg__SVGAnimatedNumberList__baseVal">baseVal</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> <a href="types.html#__svg__SVGAnimatedNumberList__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedNumberList__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a>)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.
</div></dd>
<dt id="__svg__SVGAnimatedNumberList__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a>)</span></dt><dd class="attribute"><div>
 A <a href="#ReadOnlyNumberList">read only</a> <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> representing the current animated value of
 the given attribute.  If the given attribute is not currently being
 animated, then the <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> will have the same contents
 as <a class="idlattr" href="#__svg__SVGAnimatedNumberList__baseVal">baseVal</a>.  The object referenced by <a class="idlattr" href="#__svg__SVGAnimatedNumberList__animVal">animVal</a> will always
 be distinct from the one referenced by <a class="idlattr" href="#__svg__SVGAnimatedNumberList__baseVal">baseVal</a>, even when
 the attribute is not animated.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGLength">4.5.11 Interface SVGLength</h3>


 <p>The <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> interface corresponds to the
 <a href="types.html#DataTypeLength">&lt;length&gt;</a> basic data type.</p>

 <p id="ReadOnlyLength">An <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> object can be designated as <em>read only</em>,
 which means that attempts to modify the object will result in an exception
 being thrown, as described below.</p>
<pre class="idl">interface <b>SVGLength</b> {

  // Length Unit Types
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_UNKNOWN">SVG_LENGTHTYPE_UNKNOWN</a> = 0;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_NUMBER">SVG_LENGTHTYPE_NUMBER</a> = 1;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_PERCENTAGE">SVG_LENGTHTYPE_PERCENTAGE</a> = 2;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_EMS">SVG_LENGTHTYPE_EMS</a> = 3;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_EXS">SVG_LENGTHTYPE_EXS</a> = 4;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_PX">SVG_LENGTHTYPE_PX</a> = 5;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_CM">SVG_LENGTHTYPE_CM</a> = 6;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_MM">SVG_LENGTHTYPE_MM</a> = 7;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_IN">SVG_LENGTHTYPE_IN</a> = 8;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_PT">SVG_LENGTHTYPE_PT</a> = 9;
  const unsigned short <a href="types.html#__svg__SVGLength__SVG_LENGTHTYPE_PC">SVG_LENGTHTYPE_PC</a> = 10;

  readonly attribute unsigned short <a href="types.html#__svg__SVGLength__unitType">unitType</a>;
           attribute float <a href="types.html#__svg__SVGLength__value">value</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
           attribute float <a href="types.html#__svg__SVGLength__valueInSpecifiedUnits">valueInSpecifiedUnits</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
           attribute DOMString <a href="types.html#__svg__SVGLength__valueAsString">valueAsString</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);

  void <a href="types.html#__svg__SVGLength__newValueSpecifiedUnits">newValueSpecifiedUnits</a>(in unsigned short unitType, in float valueInSpecifiedUnits) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  void <a href="types.html#__svg__SVGLength__convertToSpecifiedUnits">convertToSpecifiedUnits</a>(in unsigned short unitType) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="constants-header">Constants in group “Length Unit Types”:</dt><dd><dl class="constants">
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_UNKNOWN" class="constant first-child"><b>SVG_LENGTHTYPE_UNKNOWN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The unit type is not one of predefined unit types. It is invalid to
 attempt to define a new value of this type or to attempt to switch an
 existing value to this type.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_NUMBER" class="constant"><b>SVG_LENGTHTYPE_NUMBER</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 No unit type was provided (i.e., a unitless value was specified), which
 indicates a value in user units.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_PERCENTAGE" class="constant"><b>SVG_LENGTHTYPE_PERCENTAGE</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A percentage value was specified.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_EMS" class="constant"><b>SVG_LENGTHTYPE_EMS</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the em units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_EXS" class="constant"><b>SVG_LENGTHTYPE_EXS</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the ex units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_PX" class="constant"><b>SVG_LENGTHTYPE_PX</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the px units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_CM" class="constant"><b>SVG_LENGTHTYPE_CM</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the cm units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_MM" class="constant"><b>SVG_LENGTHTYPE_MM</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the mm units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_IN" class="constant"><b>SVG_LENGTHTYPE_IN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the in units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_PT" class="constant"><b>SVG_LENGTHTYPE_PT</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the pt units defined in CSS2.

</div></dd>
<dt id="__svg__SVGLength__SVG_LENGTHTYPE_PC" class="constant"><b>SVG_LENGTHTYPE_PC</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 A value was specified using the pc units defined in CSS2.

</div></dd></dl></dd><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLength__unitType" class="attribute first-child"><b>unitType</b><span class="idl-type-parenthetical"> (readonly unsigned short)</span></dt><dd class="attribute"><div>
 The type of the value as specified by one of the SVG_LENGTHTYPE_*
 constants defined on this interface.
</div></dd>
<dt id="__svg__SVGLength__value" class="attribute"><b>value</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The value as a floating point value, in user units. Setting this
 attribute will cause <a class="idlattr" href="#__svg__SVGLength__valueInSpecifiedUnits">valueInSpecifiedUnits</a> and
 <a class="idlattr" href="#__svg__SVGLength__valueAsString">valueAsString</a> to be updated automatically to reflect this setting.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the length
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLength">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLength__valueInSpecifiedUnits" class="attribute"><b>valueInSpecifiedUnits</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The value as a floating point value, in the units expressed by
 <a class="idlattr" href="#__svg__SVGLength__unitType">unitType</a>. Setting this attribute will cause <a class="idlattr" href="#__svg__SVGLength__value">value</a> and
 <a class="idlattr" href="#__svg__SVGLength__valueAsString">valueAsString</a> to be updated automatically to reflect this setting. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the length
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLength">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLength__valueAsString" class="attribute"><b>valueAsString</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 The value as a string value, in the units expressed by <a class="idlattr" href="#__svg__SVGLength__unitType">unitType</a>.
 Setting this attribute will cause <a class="idlattr" href="#__svg__SVGLength__value">value</a>, 
 <a class="idlattr" href="#__svg__SVGLength__valueInSpecifiedUnits">valueInSpecifiedUnits</a> and <a class="idlattr" href="#__svg__SVGLength__unitType">unitType</a> 
 to be updated automatically to reflect this
 setting. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code SYNTAX_ERR</dt><dd class="exception"> Raised if the assigned string cannot 
   be parsed as a valid <a href="types.html#DataTypeLength">&lt;length&gt;</a>.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the length
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLength">read only</a>.
</dd></dl></dd></dl></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLength__newValueSpecifiedUnits" class="operation first-child">void <b>newValueSpecifiedUnits</b>(in unsigned short <var>unitType</var>, in float <var>valueInSpecifiedUnits</var>)</dt><dd class="operation"><div>
 Reset the value as a number with an associated <a class="idlattr" href="#__svg__SVGLength__unitType">unitType</a>, thereby
 replacing the values for all of the attributes on the object.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned short <var>unitType</var></div> <div> The unit type for the value (e.g., <a class="idlattr" href="#__svg__SVGLength__SVG_LENGTHTYPE_MM">SVG_LENGTHTYPE_MM</a>).
</div></li><li class="parameter"><div>float <var>valueInSpecifiedUnits</var></div> <div> The new value.
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NOT_SUPPORTED_ERR</dt><dd class="exception"> Raised if unitType is SVG_LENGTHTYPE_UNKNOWN 
   or not a valid unit type constant (one of the other SVG_LENGTHTYPE_* constants 
   defined on this interface).
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the length
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLength">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLength__convertToSpecifiedUnits" class="operation">void <b>convertToSpecifiedUnits</b>(in unsigned short <var>unitType</var>)</dt><dd class="operation"><div>
 Preserve the same underlying stored value, but reset the stored unit
 identifier to the given <var>unitType</var>. Object attributes
 <a class="idlattr" href="#__svg__SVGLength__unitType">unitType</a>, <a class="idlattr" href="#__svg__SVGLength__valueInSpecifiedUnits">valueInSpecifiedUnits</a> and <a class="idlattr" href="#__svg__SVGLength__valueAsString">valueAsString</a>
 might be modified as a result of this method. For example, if the
 original value were "0.5cm" and the method was invoked to convert to
 millimeters, then the <a class="idlattr" href="#__svg__SVGLength__unitType">unitType</a> would be changed to
 <a class="idlattr" href="#__svg__SVGLength__SVG_LENGTHTYPE_MM">SVG_LENGTHTYPE_MM</a>, <a class="idlattr" href="#__svg__SVGLength__valueInSpecifiedUnits">valueInSpecifiedUnits</a> would be changed
 to the numeric value 5 and <a class="idlattr" href="#__svg__SVGLength__valueAsString">valueAsString</a> would be changed to
 "5mm".

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned short <var>unitType</var></div> <div> The unit type to switch to (e.g., <a class="idlattr" href="#__svg__SVGLength__SVG_LENGTHTYPE_MM">SVG_LENGTHTYPE_MM</a>).
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NOT_SUPPORTED_ERR</dt><dd class="exception"> Raised if unitType is SVG_LENGTHTYPE_UNKNOWN 
   or not a valid unit type constant (one of the other SVG_LENGTHTYPE_* constants 
   defined on this interface).
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the length
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLength">read only</a>.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedLength">4.5.12 Interface SVGAnimatedLength</h3>


 Used for attributes of basic type
 <a href="types.html#DataTypeLength">&lt;length&gt;</a> which can be
 animated.
<pre class="idl">interface <b>SVGAnimatedLength</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGAnimatedLength__baseVal">baseVal</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGAnimatedLength__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedLength__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a>)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.
</div></dd>
<dt id="__svg__SVGAnimatedLength__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a>)</span></dt><dd class="attribute"><div>
 A <a href="#ReadOnlyLength">read only</a> <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> representing the current animated value of
 the given attribute.  If the given attribute is not currently being
 animated, then the <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> will have the same contents
 as <a class="idlattr" href="#__svg__SVGAnimatedLength__baseVal">baseVal</a>.  The object referenced by <a class="idlattr" href="#__svg__SVGAnimatedLength__animVal">animVal</a> will always
 be distinct from the one referenced by <a class="idlattr" href="#__svg__SVGAnimatedLength__baseVal">baseVal</a>, even when
 the attribute is not animated.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGLengthList">4.5.13 Interface SVGLengthList</h3>


 <p>This interface defines a list of SVGLength objects.</p>

 <p><a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> has the same attributes and methods as other
 SVGxxxList interfaces. Implementers may consider using a single base class
 to implement the various SVGxxxList interfaces.</p>

 <p id="ReadOnlyLengthList">An <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> object can be designated as <em>read only</em>,
 which means that attempts to modify the object will result in an exception
 being thrown, as described below.</p>
<pre class="idl">interface <b>SVGLengthList</b> {

  readonly attribute unsigned long <a href="types.html#__svg__SVGLengthList__numberOfItems">numberOfItems</a>;

  void <a href="types.html#__svg__SVGLengthList__clear">clear</a>() raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGLengthList__initialize">initialize</a>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> newItem) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGLengthList__getItem">getItem</a>(in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGLengthList__insertItemBefore">insertItemBefore</a>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> newItem, in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGLengthList__replaceItem">replaceItem</a>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> newItem, in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGLengthList__removeItem">removeItem</a>(in unsigned long index) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <a href="types.html#__svg__SVGLengthList__appendItem">appendItem</a>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> newItem) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLengthList__numberOfItems" class="attribute first-child"><b>numberOfItems</b><span class="idl-type-parenthetical"> (readonly unsigned long)</span></dt><dd class="attribute"><div>
 The number of items in the list.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLengthList__clear" class="operation first-child">void <b>clear</b>()</dt><dd class="operation"><div>
 Clears all existing current items from the list, with the result being
 an empty list.

</div><dl class="operation"><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLengthList__initialize" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <b>initialize</b>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var>)</dt><dd class="operation"><div>
 Clears all existing current items from the list and re-initializes the
 list to hold the single item specified by the parameter.  If the inserted
 item is already in a list, it is removed from its previous list before
 it is inserted into this list.  The inserted item is the item itself and
 not a copy. 

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var></div> <div> The item which should become the only member of the list.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The item being inserted into the list.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLengthList__getItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <b>getItem</b>(in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Returns the specified item from the list.  The returned item is the
 item itself and not a copy.  Any changes made to the item are
 immediately reflected in the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned long <var>index</var></div> <div> The index of the item from the list which is to be
   returned.  The first item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The selected item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLengthList__insertItemBefore" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <b>insertItemBefore</b>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var>, in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Inserts a new item into the list at the specified position. The first
 item is number 0. If <var>newItem</var> is already in a list, it is
 removed from its previous list before it is inserted into this list.
 The inserted item is the item itself and not a copy. If the item is
 already in this list, note that the index of the item to insert
 before is <i>before</i> the removal of the item.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var></div> <div> The item which is to be inserted into the list.
</div></li><li class="parameter"><div>unsigned long <var>index</var></div> <div> The index of the item before which the new item is to be
   inserted. The first item is number 0.  If the index is equal to 0,
   then the new item is inserted at the front of the list. If the index
   is greater than or equal to <a class="idlattr" href="#__svg__SVGLengthList__numberOfItems">numberOfItems</a>, then the new item is
   appended to the end of the list.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLengthList__replaceItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <b>replaceItem</b>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var>, in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Replaces an existing item in the list with a new item. If
 <var>newItem</var> is already in a list, it is removed from its
 previous list before it is inserted into this list.  The inserted item
 is the item itself and not a copy.  If the item is already in this
 list, note that the index of the item to replace is <i>before</i>
 the removal of the item.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var></div> <div> The item which is to be inserted into the list.
</div></li><li class="parameter"><div>unsigned long <var>index</var></div> <div> The index of the item which is to be replaced. The first
   item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGLengthList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLengthList__removeItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <b>removeItem</b>(in unsigned long <var>index</var>)</dt><dd class="operation"><div>
 Removes an existing item from the list.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned long <var>index</var></div> <div> The index of the item which is to be removed. The first
   item is number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The removed item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code INDEX_SIZE_ERR</dt><dd class="exception"> Raised if the index number is
   greater than or equal to <a class="idlattr" href="#__svg__SVGLengthList__numberOfItems">numberOfItems</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLengthList__appendItem" class="operation"><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <b>appendItem</b>(in <a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var>)</dt><dd class="operation"><div>
 Inserts a new item at the end of the list. If <var>newItem</var> is
 already in a list, it is removed from its previous list before it is
 inserted into this list.  The inserted item is the item itself and
 not a copy.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGLength">SVGLength</a> <var>newItem</var></div> <div> The item which is to be inserted. The first item is
   number 0.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The inserted item.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the list
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyLengthList">read only</a>.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedLengthList">4.5.14 Interface SVGAnimatedLengthList</h3>


 Used for attributes of type <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> which can be animated.
<pre class="idl">interface <b>SVGAnimatedLengthList</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> <a href="types.html#__svg__SVGAnimatedLengthList__baseVal">baseVal</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> <a href="types.html#__svg__SVGAnimatedLengthList__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedLengthList__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a>)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.
</div></dd>
<dt id="__svg__SVGAnimatedLengthList__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a>)</span></dt><dd class="attribute"><div>
 A <a href="#ReadOnlyLengthList">read only</a> <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> representing the current animated value of
 the given attribute.  If the given attribute is not currently being
 animated, then the <a class="idlinterface" href="types.html#InterfaceSVGLengthList">SVGLengthList</a> will have the same contents
 as <a class="idlattr" href="#__svg__SVGAnimatedLengthList__baseVal">baseVal</a>.  The object referenced by <a class="idlattr" href="#__svg__SVGAnimatedLengthList__animVal">animVal</a> will always
 be distinct from the one referenced by <a class="idlattr" href="#__svg__SVGAnimatedLengthList__baseVal">baseVal</a>, even when
 the attribute is not animated.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGAngle">4.5.15 Interface SVGAngle</h3>


 <p>The <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> interface corresponds to the
 <a href="types.html#DataTypeAngle">&lt;angle&gt;</a> basic data type.</p>

 <p id="ReadOnlyAngle">An <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> object can be designated as <em>read only</em>,
 which means that attempts to modify the object will result in an exception
 being thrown, as described below.</p>
<pre class="idl">interface <b>SVGAngle</b> {

  // Angle Unit Types
  const unsigned short <a href="types.html#__svg__SVGAngle__SVG_ANGLETYPE_UNKNOWN">SVG_ANGLETYPE_UNKNOWN</a> = 0;
  const unsigned short <a href="types.html#__svg__SVGAngle__SVG_ANGLETYPE_UNSPECIFIED">SVG_ANGLETYPE_UNSPECIFIED</a> = 1;
  const unsigned short <a href="types.html#__svg__SVGAngle__SVG_ANGLETYPE_DEG">SVG_ANGLETYPE_DEG</a> = 2;
  const unsigned short <a href="types.html#__svg__SVGAngle__SVG_ANGLETYPE_RAD">SVG_ANGLETYPE_RAD</a> = 3;
  const unsigned short <a href="types.html#__svg__SVGAngle__SVG_ANGLETYPE_GRAD">SVG_ANGLETYPE_GRAD</a> = 4;

  readonly attribute unsigned short <a href="types.html#__svg__SVGAngle__unitType">unitType</a>;
           attribute float <a href="types.html#__svg__SVGAngle__value">value</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
           attribute float <a href="types.html#__svg__SVGAngle__valueInSpecifiedUnits">valueInSpecifiedUnits</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
           attribute DOMString <a href="types.html#__svg__SVGAngle__valueAsString">valueAsString</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);

  void <a href="types.html#__svg__SVGAngle__newValueSpecifiedUnits">newValueSpecifiedUnits</a>(in unsigned short unitType, in float valueInSpecifiedUnits) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  void <a href="types.html#__svg__SVGAngle__convertToSpecifiedUnits">convertToSpecifiedUnits</a>(in unsigned short unitType) raises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="constants-header">Constants in group “Angle Unit Types”:</dt><dd><dl class="constants">
<dt id="__svg__SVGAngle__SVG_ANGLETYPE_UNKNOWN" class="constant first-child"><b>SVG_ANGLETYPE_UNKNOWN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The unit type is not one of predefined unit types. It is invalid to
 attempt to define a new value of this type or to attempt to switch an
 existing value to this type.

</div></dd>
<dt id="__svg__SVGAngle__SVG_ANGLETYPE_UNSPECIFIED" class="constant"><b>SVG_ANGLETYPE_UNSPECIFIED</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 No unit type was provided (i.e., a unitless value was specified). For
 angles, a unitless value is treated the same as if degrees were
 specified.

</div></dd>
<dt id="__svg__SVGAngle__SVG_ANGLETYPE_DEG" class="constant"><b>SVG_ANGLETYPE_DEG</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The unit type was explicitly set to degrees.

</div></dd>
<dt id="__svg__SVGAngle__SVG_ANGLETYPE_RAD" class="constant"><b>SVG_ANGLETYPE_RAD</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The unit type is radians.

</div></dd>
<dt id="__svg__SVGAngle__SVG_ANGLETYPE_GRAD" class="constant"><b>SVG_ANGLETYPE_GRAD</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The unit type is radians.

</div></dd></dl></dd><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAngle__unitType" class="attribute first-child"><b>unitType</b><span class="idl-type-parenthetical"> (readonly unsigned short)</span></dt><dd class="attribute"><div>
 The type of the value as specified by one of the SVG_ANGLETYPE_*
 constants defined on this interface.
</div></dd>
<dt id="__svg__SVGAngle__value" class="attribute"><b>value</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The angle value as a floating point value, in degrees. Setting this
 attribute will cause <a class="idlattr" href="#__svg__SVGAngle__valueInSpecifiedUnits">valueInSpecifiedUnits</a> and
 <a class="idlattr" href="#__svg__SVGAngle__valueAsString">valueAsString</a> to be updated automatically to reflect this setting. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the angle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyAngle">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAngle__valueInSpecifiedUnits" class="attribute"><b>valueInSpecifiedUnits</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The angle value as a floating point value, in the units expressed by
 <a class="idlattr" href="#__svg__SVGAngle__unitType">unitType</a>. Setting this attribute will cause <a class="idlattr" href="#__svg__SVGAngle__value">value</a> and
 <a class="idlattr" href="#__svg__SVGAngle__valueAsString">valueAsString</a> to be updated automatically to reflect this setting. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the angle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyAngle">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAngle__valueAsString" class="attribute"><b>valueAsString</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 The angle value as a string value, in the units expressed by
 <a class="idlattr" href="#__svg__SVGAngle__unitType">unitType</a>. Setting this attribute will cause <a class="idlattr" href="#__svg__SVGAngle__value">value</a>,
 <a class="idlattr" href="#__svg__SVGAngle__valueInSpecifiedUnits">valueInSpecifiedUnits</a> and <a class="idlattr" href="#__svg__SVGAngle__unitType">unitType</a>
 to be updated automatically to reflect
 this setting. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code SYNTAX_ERR</dt><dd class="exception"> Raised if the assigned string cannot 
   be parsed as a valid <a href="types.html#DataTypeAngle">&lt;angle&gt;</a>.
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the angle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyAngle">read only</a>.
</dd></dl></dd></dl></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAngle__newValueSpecifiedUnits" class="operation first-child">void <b>newValueSpecifiedUnits</b>(in unsigned short <var>unitType</var>, in float <var>valueInSpecifiedUnits</var>)</dt><dd class="operation"><div>
 Reset the value as a number with an associated <a class="idlattr" href="#__svg__SVGAngle__unitType">unitType</a>, thereby
 replacing the values for all of the attributes on the object.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned short <var>unitType</var></div> <div> The unit type for the value (e.g., <a class="idlattr" href="#__svg__SVGAngle__SVG_ANGLETYPE_DEG">SVG_ANGLETYPE_DEG</a>).
</div></li><li class="parameter"><div>float <var>valueInSpecifiedUnits</var></div> <div> The angle value.
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NOT_SUPPORTED_ERR</dt><dd class="exception"> Raised if unitType is SVG_ANGLETYPE_UNKNOWN 
   or not a valid unit type constant (one of the other SVG_ANGLETYPE_* constants 
   defined on this interface).
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the angle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyAngle">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGAngle__convertToSpecifiedUnits" class="operation">void <b>convertToSpecifiedUnits</b>(in unsigned short <var>unitType</var>)</dt><dd class="operation"><div>
 Preserve the same underlying stored value, but reset the stored unit
 identifier to the given <var>unitType</var>. Object attributes
 <a class="idlattr" href="#__svg__SVGAngle__unitType">unitType</a>, <a class="idlattr" href="#__svg__SVGAngle__valueInSpecifiedUnits">valueInSpecifiedUnits</a> and <a class="idlattr" href="#__svg__SVGAngle__valueAsString">valueAsString</a>
 might be modified as a result of this method. 

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned short <var>unitType</var></div> <div> The unit type to switch to (e.g., <a class="idlattr" href="#__svg__SVGAngle__SVG_ANGLETYPE_DEG">SVG_ANGLETYPE_DEG</a>).
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NOT_SUPPORTED_ERR</dt><dd class="exception"> Raised if unitType is SVG_ANGLETYPE_UNKNOWN 
   or not a valid unit type constant (one of the other SVG_ANGLETYPE_* constants 
   defined on this interface).
</dd><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the angle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyAngle">read only</a>.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedAngle">4.5.16 Interface SVGAnimatedAngle</h3>


 Used for attributes of basic data type <a href="types.html#DataTypeAngle">&lt;angle&gt;</a>
 that can be animated.
<pre class="idl">interface <b>SVGAnimatedAngle</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> <a href="types.html#__svg__SVGAnimatedAngle__baseVal">baseVal</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> <a href="types.html#__svg__SVGAnimatedAngle__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedAngle__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a>)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.
</div></dd>
<dt id="__svg__SVGAnimatedAngle__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a>)</span></dt><dd class="attribute"><div>
 A <a href="#ReadOnlyAngle">read only</a> <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> representing the current animated value of
 the given attribute.  If the given attribute is not currently being
 animated, then the <a class="idlinterface" href="types.html#InterfaceSVGAngle">SVGAngle</a> will have the same contents
 as <a class="idlattr" href="#__svg__SVGAnimatedAngle__baseVal">baseVal</a>.  The object referenced by <a class="idlattr" href="#__svg__SVGAnimatedAngle__animVal">animVal</a> will always
 be distinct from the one referenced by <a class="idlattr" href="#__svg__SVGAnimatedAngle__baseVal">baseVal</a>, even when
 the attribute is not animated.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGColor">4.5.17 Interface SVGColor</h3>


 <p>The <a class="idlinterface" href="types.html#InterfaceSVGColor">SVGColor</a> interface corresponds to color value definition for
 properties <a href="pservers.html#StopColorProperty"><span class="prop-name">‘stop-color’</span></a>, <a href="filters.html#FloodColorProperty"><span class="prop-name">‘flood-color’</span></a> and
 <a href="filters.html#LightingColorProperty"><span class="prop-name">‘lighting-color’</span></a> and is a base class for interface <a class="idlinterface" href="painting.html#InterfaceSVGPaint">SVGPaint</a>.
 It incorporates SVG's extended notion of color, which incorporates
 ICC-based color specifications.</p>

 <p>Interface <a class="idlinterface" href="types.html#InterfaceSVGColor">SVGColor</a> does <em>not</em> correspond to the
 <a href="types.html#DataTypeColor">&lt;color&gt;</a> basic data type. For
 the <a href="types.html#DataTypeColor">&lt;color&gt;</a> basic data type,
 the applicable DOM interfaces are defined in
 <a href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/">DOM Level 2 Style</a>;
 in particular, see the <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-RGBColor">RGBColor</a> interface
 ([<a href="refs.html#ref-DOM2STYLE">DOM2STYLE</a>], section 2.2).</p>

 <p>Note: The <a class="idlinterface" href="types.html#InterfaceSVGColor">SVGColor</a> interface is deprecated, and may be dropped 
 from future versions of the SVG specification.</p>
<pre class="idl">interface <b>SVGColor</b> : <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSValue">CSSValue</a> {

  // Color Types
  const unsigned short <a href="types.html#__svg__SVGColor__SVG_COLORTYPE_UNKNOWN">SVG_COLORTYPE_UNKNOWN</a> = 0;
  const unsigned short <a href="types.html#__svg__SVGColor__SVG_COLORTYPE_RGBCOLOR">SVG_COLORTYPE_RGBCOLOR</a> = 1;
  const unsigned short <a href="types.html#__svg__SVGColor__SVG_COLORTYPE_RGBCOLOR_ICCCOLOR">SVG_COLORTYPE_RGBCOLOR_ICCCOLOR</a> = 2;
  const unsigned short <a href="types.html#__svg__SVGColor__SVG_COLORTYPE_CURRENTCOLOR">SVG_COLORTYPE_CURRENTCOLOR</a> = 3;

  readonly attribute unsigned short <a href="types.html#__svg__SVGColor__colorType">colorType</a>;
  readonly attribute <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-RGBColor">RGBColor</a> <a href="types.html#__svg__SVGColor__rgbColor">rgbColor</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGICCColor">SVGICCColor</a> <a href="types.html#__svg__SVGColor__iccColor">iccColor</a>;

  void <a href="types.html#__svg__SVGColor__setRGBColor">setRGBColor</a>(in DOMString rgbColor) raises(<a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>);
  void <a href="types.html#__svg__SVGColor__setRGBColorICCColor">setRGBColorICCColor</a>(in DOMString rgbColor, in DOMString iccColor) raises(<a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>);
  void <a href="types.html#__svg__SVGColor__setColor">setColor</a>(in unsigned short colorType, in DOMString rgbColor, in DOMString iccColor) raises(<a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>);
};</pre><dl class="interface"><dt class="constants-header">Constants in group “Color Types”:</dt><dd><dl class="constants">
<dt id="__svg__SVGColor__SVG_COLORTYPE_UNKNOWN" class="constant first-child"><b>SVG_COLORTYPE_UNKNOWN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The color type is not one of predefined types. It is invalid to attempt
 to define a new value of this type or to attempt to switch an existing
 value to this type.

</div></dd>
<dt id="__svg__SVGColor__SVG_COLORTYPE_RGBCOLOR" class="constant"><b>SVG_COLORTYPE_RGBCOLOR</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 An sRGB color has been specified without an alternative ICC color
 specification.

</div></dd>
<dt id="__svg__SVGColor__SVG_COLORTYPE_RGBCOLOR_ICCCOLOR" class="constant"><b>SVG_COLORTYPE_RGBCOLOR_ICCCOLOR</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 An sRGB color has been specified along with an alternative ICC color
 specification.

</div></dd>
<dt id="__svg__SVGColor__SVG_COLORTYPE_CURRENTCOLOR" class="constant"><b>SVG_COLORTYPE_CURRENTCOLOR</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to when keyword <span class="attr-value">currentColor</span>
 has been specified.

</div></dd></dl></dd><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGColor__colorType" class="attribute first-child"><b>colorType</b><span class="idl-type-parenthetical"> (readonly unsigned short)</span></dt><dd class="attribute"><div>
 The type of the value as specified by one of the SVG_COLORTYPE_*
 constants defined on this interface.
</div></dd>
<dt id="__svg__SVGColor__rgbColor" class="attribute"><b>rgbColor</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-RGBColor">RGBColor</a>)</span></dt><dd class="attribute"><div>
 The color specified in the sRGB color space.
</div></dd>
<dt id="__svg__SVGColor__iccColor" class="attribute"><b>iccColor</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGICCColor">SVGICCColor</a>)</span></dt><dd class="attribute"><div>
 The alternate ICC color specification.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGColor__setRGBColor" class="operation first-child">void <b>setRGBColor</b>(in DOMString <var>rgbColor</var>)</dt><dd class="operation"><div>
 Modifies the color value to be the specified sRGB color without an
 alternate ICC color specification.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>rgbColor</var></div> <div> A string that matches <a href="types.html#DataTypeColor">&lt;color&gt;</a>,
   which specifies the new sRGB color value.
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>, code SVG_INVALID_VALUE_ERR</dt><dd class="exception"> Raised if <var>rgbColor</var>
   does not match <a href="types.html#DataTypeColor">&lt;color&gt;</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGColor__setRGBColorICCColor" class="operation">void <b>setRGBColorICCColor</b>(in DOMString <var>rgbColor</var>, in DOMString <var>iccColor</var>)</dt><dd class="operation"><div>
 Modifies the color value to be the specified sRGB color with an
 alternate ICC color specification.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>rgbColor</var></div> <div> A string that matches <a href="types.html#DataTypeColor">&lt;color&gt;</a>,
   which specifies the new sRGB color value.
</div></li><li class="parameter"><div>DOMString <var>iccColor</var></div> <div> A string that matches <a href="types.html#DataTypeICCColor">&lt;icccolor&gt;</a>,
   which specifies the alternate ICC color specification.
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>, code SVG_INVALID_VALUE_ERR</dt><dd class="exception"> Raised if <var>rgbColor</var>
   does not match <a href="types.html#DataTypeColor">&lt;color&gt;</a>
   or if <var>iccColor</var> does not match
   <a href="types.html#DataTypeICCColor">&lt;icccolor&gt;</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGColor__setColor" class="operation">void <b>setColor</b>(in unsigned short <var>colorType</var>, in DOMString <var>rgbColor</var>, in DOMString <var>iccColor</var>)</dt><dd class="operation"><div>
 Sets the color value as specified by the parameters. If
 <var>colorType</var> requires an <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-RGBColor">RGBColor</a>, then
 <var>rgbColor</var> must be a string that matches
 <a href="types.html#DataTypeColor">&lt;color&gt;</a>;
 otherwise, <var>rgbColor</var>. must be null. If <var>colorType</var>
 requires an <a class="idlinterface" href="types.html#InterfaceSVGICCColor">SVGICCColor</a>, then <var>iccColor</var> must be a string
 that matches <a href="types.html#DataTypeICCColor">&lt;icccolor&gt;</a>;
 otherwise, <var>iccColor</var> must be null. 

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>unsigned short <var>colorType</var></div> <div> One of the defined constants for <a class="idlattr" href="#__svg__SVGColor__colorType">colorType</a>.
</div></li><li class="parameter"><div>DOMString <var>rgbColor</var></div> <div> The specification of an sRGB color, or null.
</div></li><li class="parameter"><div>DOMString <var>iccColor</var></div> <div> The specification of an ICC color, or null.
</div></li></ol></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>, code SVG_INVALID_VALUE_ERR</dt><dd class="exception"> Raised if one of the
   parameters has an invalid value.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGICCColor">4.5.18 Interface SVGICCColor</h3>


 <p>The <a class="idlinterface" href="types.html#InterfaceSVGICCColor">SVGICCColor</a> interface expresses an ICC-based color
 specification.</p>

 <p>Note: The <a class="idlinterface" href="types.html#InterfaceSVGICCColor">SVGICCColor</a> interface is deprecated, and may be dropped 
 from future versions of the SVG specification.</p>
<pre class="idl">interface <b>SVGICCColor</b> {
           attribute DOMString <a href="types.html#__svg__SVGICCColor__colorProfile">colorProfile</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a> <a href="types.html#__svg__SVGICCColor__colors">colors</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGICCColor__colorProfile" class="attribute first-child"><b>colorProfile</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 The name of the color profile, which is the first parameter of an ICC
 color specification.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an attempt
   to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGICCColor__colors" class="attribute"><b>colors</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGNumberList">SVGNumberList</a>)</span></dt><dd class="attribute"><div>
 The list of color values that define this ICC color. Each color value
 is an arbitrary floating point number.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGRect">4.5.19 Interface SVGRect</h3>


 <p>Represents rectangular geometry.  Rectangles are defined as consisting
 of a (x,y) coordinate pair identifying a minimum X value, a minimum Y
 value, and a width and height, which are usually constrained to be
 non-negative.</p>

 <p id="ReadOnlyRect">An <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> object can be designated as <em>read only</em>,
 which means that attempts to modify the object will result in an exception
 being thrown, as described below.</p>
<pre class="idl">interface <b>SVGRect</b> {
  attribute float <a href="types.html#__svg__SVGRect__x">x</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  attribute float <a href="types.html#__svg__SVGRect__y">y</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  attribute float <a href="types.html#__svg__SVGRect__width">width</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  attribute float <a href="types.html#__svg__SVGRect__height">height</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGRect__x" class="attribute first-child"><b>x</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The <var>x</var> coordinate of the rectangle, in user units. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the rectangle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyRect">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGRect__y" class="attribute"><b>y</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The <var>y</var> coordinate of the rectangle, in user units. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the rectangle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyRect">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGRect__width" class="attribute"><b>width</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The <var>width</var> coordinate of the rectangle, in user units. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the rectangle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyRect">read only</a>.
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGRect__height" class="attribute"><b>height</b><span class="idl-type-parenthetical"> (float)</span></dt><dd class="attribute"><div>
 The <var>height</var> coordinate of the rectangle, in user units. 

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised when the rectangle
   corresponds to a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a> or when the object itself is
   <a href="#ReadOnlyRect">read only</a>.
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGAnimatedRect">4.5.20 Interface SVGAnimatedRect</h3>


 Used for attributes of type <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> which can be animated.
<pre class="idl">interface <b>SVGAnimatedRect</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> <a href="types.html#__svg__SVGAnimatedRect__baseVal">baseVal</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> <a href="types.html#__svg__SVGAnimatedRect__animVal">animVal</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGAnimatedRect__baseVal" class="attribute first-child"><b>baseVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a>)</span></dt><dd class="attribute"><div>
 The base value of the given attribute before applying any animations.
</div></dd>
<dt id="__svg__SVGAnimatedRect__animVal" class="attribute"><b>animVal</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a>)</span></dt><dd class="attribute"><div>
 A <a href="#ReadOnlyRect">read only</a> <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> representing the current animated value of
 the given attribute.  If the given attribute is not currently being
 animated, then the <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> will have the same contents
 as <a class="idlattr" href="#__svg__SVGAnimatedRect__baseVal">baseVal</a>.  The object referenced by <a class="idlattr" href="#__svg__SVGAnimatedRect__animVal">animVal</a> will always
 be distinct from the one referenced by <a class="idlattr" href="#__svg__SVGAnimatedRect__baseVal">baseVal</a>, even when
 the attribute is not animated.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGUnitTypes">4.5.21 Interface SVGUnitTypes</h3>


 The <a class="idlinterface" href="types.html#InterfaceSVGUnitTypes">SVGUnitTypes</a> interface defines a commonly used set of constants
 and is a base interface used by <a class="idlinterface" href="pservers.html#InterfaceSVGGradientElement">SVGGradientElement</a>,
 <a class="idlinterface" href="pservers.html#InterfaceSVGPatternElement">SVGPatternElement</a>, <a class="idlinterface" href="masking.html#InterfaceSVGClipPathElement">SVGClipPathElement</a>, <a class="idlinterface" href="masking.html#InterfaceSVGMaskElement">SVGMaskElement</a>
 and <a class="idlinterface" href="filters.html#InterfaceSVGFilterElement">SVGFilterElement</a>.
<pre class="idl">interface <b>SVGUnitTypes</b> {
  // Unit Types
  const unsigned short <a href="types.html#__svg__SVGUnitTypes__SVG_UNIT_TYPE_UNKNOWN">SVG_UNIT_TYPE_UNKNOWN</a> = 0;
  const unsigned short <a href="types.html#__svg__SVGUnitTypes__SVG_UNIT_TYPE_USERSPACEONUSE">SVG_UNIT_TYPE_USERSPACEONUSE</a> = 1;
  const unsigned short <a href="types.html#__svg__SVGUnitTypes__SVG_UNIT_TYPE_OBJECTBOUNDINGBOX">SVG_UNIT_TYPE_OBJECTBOUNDINGBOX</a> = 2;
};</pre><dl class="interface"><dt class="constants-header">Constants in group “Unit Types”:</dt><dd><dl class="constants">
<dt id="__svg__SVGUnitTypes__SVG_UNIT_TYPE_UNKNOWN" class="constant first-child"><b>SVG_UNIT_TYPE_UNKNOWN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The type is not one of predefined types. It is invalid to attempt to
 define a new value of this type or to attempt to switch an existing
 value to this type.

</div></dd>
<dt id="__svg__SVGUnitTypes__SVG_UNIT_TYPE_USERSPACEONUSE" class="constant"><b>SVG_UNIT_TYPE_USERSPACEONUSE</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to value <span class="attr-value">'userSpaceOnUse'</span>.

</div></dd>
<dt id="__svg__SVGUnitTypes__SVG_UNIT_TYPE_OBJECTBOUNDINGBOX" class="constant"><b>SVG_UNIT_TYPE_OBJECTBOUNDINGBOX</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to value <span class="attr-value">'objectBoundingBox'</span>.

</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGStylable">4.5.22 Interface SVGStylable</h3>


 The <a class="idlinterface" href="types.html#InterfaceSVGStylable">SVGStylable</a> interface is implemented on all objects
 corresponding to SVG elements that can have <a href="styling.html#StyleAttribute"><span class="attr-name">‘style’</span></a>,
 <a href="styling.html#ClassAttribute"><span class="attr-name">‘class’</span></a> and <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attributes</span></a> specified on them.  It
 is thus an ancestor interface for many of the interfaces defined in this
 specification.
<pre class="idl">interface <b>SVGStylable</b> {

  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGAnimatedString">SVGAnimatedString</a> <a href="types.html#__svg__SVGStylable__className">className</a>;
  readonly attribute <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSStyleDeclaration">CSSStyleDeclaration</a> <a href="types.html#__svg__SVGStylable__style">style</a>;

  <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSValue">CSSValue</a> <a href="types.html#__svg__SVGStylable__getPresentationAttribute">getPresentationAttribute</a>(in DOMString name);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGStylable__className" class="attribute first-child"><b>className</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGAnimatedString">SVGAnimatedString</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="styling.html#ClassAttribute"><span class="attr-name">‘class’</span></a> on the given element.
</div></dd>
<dt id="__svg__SVGStylable__style" class="attribute"><b>style</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSStyleDeclaration">CSSStyleDeclaration</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="styling.html#StyleAttribute"><span class="attr-name">‘style’</span></a> on the given element. If the
 user agent does not support <a href="styling.html#StylingWithCSS">styling
   with CSS</a>, then this attribute must always have the value of null.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGStylable__getPresentationAttribute" class="operation first-child"><a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSValue">CSSValue</a> <b>getPresentationAttribute</b>(in DOMString <var>name</var>)</dt><dd class="operation"><div>
 Returns the base (i.e., static) value of a given <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a> as an object of type <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSValue">CSSValue</a>. The returned object
 is live; changes to the objects represent immediate changes to the
 objects to which the <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSValue">CSSValue</a> is attached.

 <p>Note: The <code>getPresentationAttribute</code> method is deprecated, 
 and may be dropped from future versions of the SVG specification.</p>

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>name</var></div> <div> The name of the presentation attribute whose value is to be
   returned.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> The static/base value of the given <a href="intro.html#TermPresentationAttribute"><span class="svg-term">presentation attribute</span></a>
   as a <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSValue">CSSValue</a>, or null if the given attribute does not have a
   specified value.
</div></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGLocatable">4.5.23 Interface SVGLocatable</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGLocatable">SVGLocatable</a> is for all elements which either have a
 <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute or don't have a <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute
 but whose content can have a bounding box in current user space.
<pre class="idl">interface <b>SVGLocatable</b> {

  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> <a href="types.html#__svg__SVGLocatable__nearestViewportElement">nearestViewportElement</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> <a href="types.html#__svg__SVGLocatable__farthestViewportElement">farthestViewportElement</a>;

  <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> <a href="types.html#__svg__SVGLocatable__getBBox">getBBox</a>();
  <a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> <a href="types.html#__svg__SVGLocatable__getCTM">getCTM</a>();
  <a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> <a href="types.html#__svg__SVGLocatable__getScreenCTM">getScreenCTM</a>();
  <a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> <a href="types.html#__svg__SVGLocatable__getTransformToElement">getTransformToElement</a>(in <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> element) raises(<a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLocatable__nearestViewportElement" class="attribute first-child"><b>nearestViewportElement</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a>)</span></dt><dd class="attribute"><div>
 The element which established the current viewport. Often, the nearest
 ancestor <a href="struct.html#SVGElement"><span class="element-name">‘svg’</span></a> element. Null if the current element is the
 <a href="intro.html#TermOutermostSVGElement"><span class="svg-term">outermost svg element</span></a>.
</div></dd>
<dt id="__svg__SVGLocatable__farthestViewportElement" class="attribute"><b>farthestViewportElement</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a>)</span></dt><dd class="attribute"><div>
 The farthest ancestor <a href="struct.html#SVGElement"><span class="element-name">‘svg’</span></a> element. Null if the current element
 is the <a href="intro.html#TermOutermostSVGElement"><span class="svg-term">outermost svg element</span></a>.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLocatable__getBBox" class="operation first-child"><a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> <b>getBBox</b>()</dt><dd class="operation"><div>
 Returns the tight bounding box in current user space (i.e., after
 application of the <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute, if any) on the
 geometry of all contained graphics elements, exclusive of stroking, clipping, masking and
 filter effects). Note that getBBox must return the actual bounding box
 at the time the method was called, even in case the element has not
 yet been rendered. 

</div><dl class="operation"><dt class="returns-header">Returns</dt><dd><div> An <a class="idlinterface" href="types.html#InterfaceSVGRect">SVGRect</a> object that defines the bounding box.
</div></dd></dl></dd>
<dt id="__svg__SVGLocatable__getCTM" class="operation"><a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> <b>getCTM</b>()</dt><dd class="operation"><div>
 Returns the transformation matrix from current user units (i.e., after
 application of the <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute, if any) to the viewport
 coordinate system for the <a class="idlattr" href="#__svg__SVGLocatable__nearestViewportElement">nearestViewportElement</a>. 

</div><dl class="operation"><dt class="returns-header">Returns</dt><dd><div> An <a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> object that defines the CTM.
</div></dd></dl></dd>
<dt id="__svg__SVGLocatable__getScreenCTM" class="operation"><a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> <b>getScreenCTM</b>()</dt><dd class="operation"><div>
 Returns the transformation matrix from current user units (i.e., after
 application of the <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute, if any) to the parent
 user agent's notice of a "pixel". For display devices, ideally this
 represents a physical screen pixel. For other devices or environments
 where physical pixel sizes are not known, then an algorithm similar to
 the CSS2 definition of a "pixel" can be used instead.  Note that null 
 is returned if this element is not hooked into the document tree. This 
 method would have been more aptly named as <code>getClientCTM</code>, 
 but the name <code>getScreenCTM</code> is kept for historical reasons.

</div><dl class="operation"><dt class="returns-header">Returns</dt><dd><div> An <a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> object that defines the given
   transformation matrix.
</div></dd></dl></dd>
<dt id="__svg__SVGLocatable__getTransformToElement" class="operation"><a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> <b>getTransformToElement</b>(in <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> <var>element</var>)</dt><dd class="operation"><div>
 Returns the transformation matrix from the user coordinate system on the
 current element (after application of the <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute,
 if any) to the user coordinate system on parameter <var>element</var>
 (after application of its <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> attribute, if any). 

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div><a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> <var>element</var></div> <div> The target element.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> An <a class="idlinterface" href="coords.html#InterfaceSVGMatrix">SVGMatrix</a> object that defines the transformation.
</div></dd><dt class="exceptions-header">Exceptions</dt><dd><dl class="exceptions"><dt class="exception first-child"><a class="idlinterface" href="svgdom.html#ExceptionSVGException">SVGException</a>, code SVG_MATRIX_NOT_INVERTABLE</dt><dd class="exception"> Raised if the currently
   defined transformation matrices make it impossible to compute the
   given matrix (e.g., because one of the transformations is singular). 
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGTransformable">4.5.24 Interface SVGTransformable</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGTransformable">SVGTransformable</a> contains properties and methods that
 apply to all elements which have attribute <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a>.
<pre class="idl">interface <b>SVGTransformable</b> : <a class="idlinterface" href="types.html#InterfaceSVGLocatable">SVGLocatable</a> {
  readonly attribute <a class="idlinterface" href="coords.html#InterfaceSVGAnimatedTransformList">SVGAnimatedTransformList</a> <a href="types.html#__svg__SVGTransformable__transform">transform</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGTransformable__transform" class="attribute first-child"><b>transform</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="coords.html#InterfaceSVGAnimatedTransformList">SVGAnimatedTransformList</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="coords.html#TransformAttribute"><span class="attr-name">‘transform’</span></a> on the given element.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGTests">4.5.25 Interface SVGTests</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGTests">SVGTests</a> defines an interface which applies to all
 elements which have attributes <a href="struct.html#RequiredFeaturesAttribute"><span class="attr-name">‘requiredFeatures’</span></a>,
 <a href="struct.html#RequiredExtensionsAttribute"><span class="attr-name">‘requiredExtensions’</span></a> and <a href="struct.html#SystemLanguageAttribute"><span class="attr-name">‘systemLanguage’</span></a>.
<pre class="idl">interface <b>SVGTests</b> {

  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a> <a href="types.html#__svg__SVGTests__requiredFeatures">requiredFeatures</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a> <a href="types.html#__svg__SVGTests__requiredExtensions">requiredExtensions</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a> <a href="types.html#__svg__SVGTests__systemLanguage">systemLanguage</a>;

  boolean <a href="types.html#__svg__SVGTests__hasExtension">hasExtension</a>(in DOMString extension);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGTests__requiredFeatures" class="attribute first-child"><b>requiredFeatures</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#RequiredFeaturesAttribute"><span class="attr-name">‘requiredFeatures’</span></a> on the given element.
</div></dd>
<dt id="__svg__SVGTests__requiredExtensions" class="attribute"><b>requiredExtensions</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#RequiredExtensionsAttribute"><span class="attr-name">‘requiredExtensions’</span></a> on the given element.
</div></dd>
<dt id="__svg__SVGTests__systemLanguage" class="attribute"><b>systemLanguage</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGStringList">SVGStringList</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#SystemLanguageAttribute"><span class="attr-name">‘systemLanguage’</span></a> on the given element.
</div></dd></dl></dd><dt class="operations-header">Operations:</dt><dd><dl class="attributes">
<dt id="__svg__SVGTests__hasExtension" class="operation first-child">boolean <b>hasExtension</b>(in DOMString <var>extension</var>)</dt><dd class="operation"><div>
 Returns true if the user agent supports the given extension, specified
 by a URI.

</div><dl class="operation"><dt class="parameters-header">Parameters</dt><dd><ol class="parameters"><li class="parameter first-child"><div>DOMString <var>extension</var></div> <div> The name of the extension, expressed as a URI.
</div></li></ol></dd><dt class="returns-header">Returns</dt><dd><div> True or false, depending on whether the given extension is
   supported.
</div></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGLangSpace">4.5.26 Interface SVGLangSpace</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGLangSpace">SVGLangSpace</a> defines an interface which applies to all
 elements which have attributes <a href="struct.html#XMLLangAttribute"><span class="attr-name">‘xml:lang’</span></a> and <a href="struct.html#XMLSpaceAttribute"><span class="attr-name">‘xml:space’</span></a>.
<pre class="idl">interface <b>SVGLangSpace</b> {
  attribute DOMString <a href="types.html#__svg__SVGLangSpace__xmllang">xmllang</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
  attribute DOMString <a href="types.html#__svg__SVGLangSpace__xmlspace">xmlspace</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGLangSpace__xmllang" class="attribute first-child"><b>xmllang</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#XMLLangAttribute"><span class="attr-name">‘xml:lang’</span></a> on the given element.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>. 
</dd></dl></dd></dl></dd>
<dt id="__svg__SVGLangSpace__xmlspace" class="attribute"><b>xmlspace</b><span class="idl-type-parenthetical"> (DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#XMLSpaceAttribute"><span class="attr-name">‘xml:space’</span></a> on the given element.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>. 
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGExternalResourcesRequired">4.5.27 Interface SVGExternalResourcesRequired</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGExternalResourcesRequired">SVGExternalResourcesRequired</a> defines an interface which
 applies to all elements where this element or one of its descendants can
 reference an external resource.
<pre class="idl">interface <b>SVGExternalResourcesRequired</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGAnimatedBoolean">SVGAnimatedBoolean</a> <a href="types.html#__svg__SVGExternalResourcesRequired__externalResourcesRequired">externalResourcesRequired</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGExternalResourcesRequired__externalResourcesRequired" class="attribute first-child"><b>externalResourcesRequired</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGAnimatedBoolean">SVGAnimatedBoolean</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="struct.html#ExternalResourcesRequiredAttribute"><span class="attr-name">‘externalResourcesRequired’</span></a> on the given
 element. Note that the SVG DOM defines the attribute
 <a href="struct.html#ExternalResourcesRequiredAttribute"><span class="attr-name">‘externalResourcesRequired’</span></a> as being of type
 <a class="idlinterface" href="types.html#InterfaceSVGAnimatedBoolean">SVGAnimatedBoolean</a>, whereas the SVG language definition says that
 <a href="struct.html#ExternalResourcesRequiredAttribute"><span class="attr-name">‘externalResourcesRequired’</span></a> is not animated. Because the SVG
 language definition states that <a href="struct.html#ExternalResourcesRequiredAttribute"><span class="attr-name">‘externalResourcesRequired’</span></a>
 cannot be animated, the <a class="idlattr" href="types.html#__svg__SVGAnimatedBoolean__animVal">animVal</a> will always be
 the same as the <a class="idlattr" href="types.html#__svg__SVGAnimatedBoolean__baseVal">baseVal</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGFitToViewBox">4.5.28 Interface SVGFitToViewBox</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGFitToViewBox">SVGFitToViewBox</a> defines DOM attributes that apply to
 elements which have XML attributes <a href="coords.html#ViewBoxAttribute"><span class="attr-name">‘viewBox’</span></a> and
 <a href="coords.html#PreserveAspectRatioAttribute"><span class="attr-name">‘preserveAspectRatio’</span></a>.
<pre class="idl">interface <b>SVGFitToViewBox</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGAnimatedRect">SVGAnimatedRect</a> <a href="types.html#__svg__SVGFitToViewBox__viewBox">viewBox</a>;
  readonly attribute <a class="idlinterface" href="coords.html#InterfaceSVGAnimatedPreserveAspectRatio">SVGAnimatedPreserveAspectRatio</a> <a href="types.html#__svg__SVGFitToViewBox__preserveAspectRatio">preserveAspectRatio</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGFitToViewBox__viewBox" class="attribute first-child"><b>viewBox</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGAnimatedRect">SVGAnimatedRect</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="coords.html#ViewBoxAttribute"><span class="attr-name">‘viewBox’</span></a> on the given element.
</div></dd>
<dt id="__svg__SVGFitToViewBox__preserveAspectRatio" class="attribute"><b>preserveAspectRatio</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="coords.html#InterfaceSVGAnimatedPreserveAspectRatio">SVGAnimatedPreserveAspectRatio</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="coords.html#PreserveAspectRatioAttribute"><span class="attr-name">‘preserveAspectRatio’</span></a> on the given element.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGZoomAndPan">4.5.29 Interface SVGZoomAndPan</h3>


 The <a class="idlinterface" href="types.html#InterfaceSVGZoomAndPan">SVGZoomAndPan</a> interface defines attribute <a class="idlattr" href="#__svg__SVGZoomAndPan__zoomAndPan">zoomAndPan</a> and
 associated constants. 
<pre class="idl">interface <b>SVGZoomAndPan</b> {

  // Zoom and Pan Types
  const unsigned short <a href="types.html#__svg__SVGZoomAndPan__SVG_ZOOMANDPAN_UNKNOWN">SVG_ZOOMANDPAN_UNKNOWN</a> = 0;
  const unsigned short <a href="types.html#__svg__SVGZoomAndPan__SVG_ZOOMANDPAN_DISABLE">SVG_ZOOMANDPAN_DISABLE</a> = 1;
  const unsigned short <a href="types.html#__svg__SVGZoomAndPan__SVG_ZOOMANDPAN_MAGNIFY">SVG_ZOOMANDPAN_MAGNIFY</a> = 2;

  attribute unsigned short <a href="types.html#__svg__SVGZoomAndPan__zoomAndPan">zoomAndPan</a> setraises(<a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>);
};</pre><dl class="interface"><dt class="constants-header">Constants in group “Zoom and Pan Types”:</dt><dd><dl class="constants">
<dt id="__svg__SVGZoomAndPan__SVG_ZOOMANDPAN_UNKNOWN" class="constant first-child"><b>SVG_ZOOMANDPAN_UNKNOWN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The enumeration was set to a value that is not one of predefined types.
 It is invalid to attempt to define a new value of this type or to
 attempt to switch an existing value to this type.

</div></dd>
<dt id="__svg__SVGZoomAndPan__SVG_ZOOMANDPAN_DISABLE" class="constant"><b>SVG_ZOOMANDPAN_DISABLE</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to value <span class="attr-value">'disable'</span>.

</div></dd>
<dt id="__svg__SVGZoomAndPan__SVG_ZOOMANDPAN_MAGNIFY" class="constant"><b>SVG_ZOOMANDPAN_MAGNIFY</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to value <span class="attr-value">'magnify'</span>.

</div></dd></dl></dd><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGZoomAndPan__zoomAndPan" class="attribute first-child"><b>zoomAndPan</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <a href="interact.html#ZoomAndPanAttribute"><span class="attr-name">‘zoomAndPan’</span></a> on the given element. The
 value must be one of the SVG_ZOOMANDPAN_* constants defined on this
 interface.

</div><dl class="attribute"><dt class="exceptions-header">Exceptions on setting</dt><dd><dl class="exceptions"><dt class="exception"><a class="idlinterface" href="http://www.w3.org/TR/DOM-Level-2-Core/core.html#ID-17189187">DOMException</a>, code NO_MODIFICATION_ALLOWED_ERR</dt><dd class="exception"> Raised on an
   attempt to change the value of a <a href="svgdom.html#ReadOnlyNodes">read only attribute</a>. 
</dd></dl></dd></dl></dd></dl></dd></dl>

<h3 id="InterfaceSVGViewSpec">4.5.30 Interface SVGViewSpec</h3>


 The interface corresponds to an SVG View Specification.
<pre class="idl">interface <b>SVGViewSpec</b> : <a class="idlinterface" href="types.html#InterfaceSVGZoomAndPan">SVGZoomAndPan</a>,
                        <a class="idlinterface" href="types.html#InterfaceSVGFitToViewBox">SVGFitToViewBox</a> {
  readonly attribute <a class="idlinterface" href="coords.html#InterfaceSVGTransformList">SVGTransformList</a> <a href="types.html#__svg__SVGViewSpec__transform">transform</a>;
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a> <a href="types.html#__svg__SVGViewSpec__viewTarget">viewTarget</a>;
  readonly attribute DOMString <a href="types.html#__svg__SVGViewSpec__viewBoxString">viewBoxString</a>;
  readonly attribute DOMString <a href="types.html#__svg__SVGViewSpec__preserveAspectRatioString">preserveAspectRatioString</a>;
  readonly attribute DOMString <a href="types.html#__svg__SVGViewSpec__transformString">transformString</a>;
  readonly attribute DOMString <a href="types.html#__svg__SVGViewSpec__viewTargetString">viewTargetString</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGViewSpec__transform" class="attribute first-child"><b>transform</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="coords.html#InterfaceSVGTransformList">SVGTransformList</a>)</span></dt><dd class="attribute"><div>
 Corresponds to the transform setting on the SVG View Specification.
</div></dd>
<dt id="__svg__SVGViewSpec__viewTarget" class="attribute"><b>viewTarget</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGElement">SVGElement</a>)</span></dt><dd class="attribute"><div>
 Corresponds to the viewTarget setting on the SVG View Specification.
</div></dd>
<dt id="__svg__SVGViewSpec__viewBoxString" class="attribute"><b>viewBoxString</b><span class="idl-type-parenthetical"> (readonly DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to the viewBox setting on the SVG View Specification.
</div></dd>
<dt id="__svg__SVGViewSpec__preserveAspectRatioString" class="attribute"><b>preserveAspectRatioString</b><span class="idl-type-parenthetical"> (readonly DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to the preserveAspectRatio setting on the SVG View Specification.
</div></dd>
<dt id="__svg__SVGViewSpec__transformString" class="attribute"><b>transformString</b><span class="idl-type-parenthetical"> (readonly DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to the transform setting on the SVG View Specification.
</div></dd>
<dt id="__svg__SVGViewSpec__viewTargetString" class="attribute"><b>viewTargetString</b><span class="idl-type-parenthetical"> (readonly DOMString)</span></dt><dd class="attribute"><div>
 Corresponds to the viewTarget setting on the SVG View Specification.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGURIReference">4.5.31 Interface SVGURIReference</h3>


 Interface <a class="idlinterface" href="types.html#InterfaceSVGURIReference">SVGURIReference</a> defines an interface which applies to all
 elements which have the collection of XLink attributes, such as
 <span class="attr-name">‘xlink:href’</span>, which define a URI reference.
<pre class="idl">interface <b>SVGURIReference</b> {
  readonly attribute <a class="idlinterface" href="types.html#InterfaceSVGAnimatedString">SVGAnimatedString</a> <a href="types.html#__svg__SVGURIReference__href">href</a>;
};</pre><dl class="interface"><dt class="attributes-header">Attributes:</dt><dd><dl class="attributes">
<dt id="__svg__SVGURIReference__href" class="attribute first-child"><b>href</b><span class="idl-type-parenthetical"> (readonly <a class="idlinterface" href="types.html#InterfaceSVGAnimatedString">SVGAnimatedString</a>)</span></dt><dd class="attribute"><div>
 Corresponds to attribute <span class="attr-name">‘xlink:href’</span> on
 the given element.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGCSSRule">4.5.32 Interface SVGCSSRule</h3>


 <p>SVG extends interface <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSRule">CSSRule</a> with interface <a class="idlinterface" href="types.html#InterfaceSVGCSSRule">SVGCSSRule</a>
 by adding an <a class="idlinterface" href="color.html#InterfaceSVGColorProfileRule">SVGColorProfileRule</a> rule to allow for specification of
 ICC-based color.</p>

 <p>It is likely that this extension will become part of a future version of
 CSS and DOM.</p>
<pre class="idl">interface <b>SVGCSSRule</b> : <a class="idlinterface" href="http://www.w3.org/TR/2000/REC-DOM-Level-2-Style-20001113/css.html#CSS-CSSRule">CSSRule</a> {
  const unsigned short <a href="types.html#__svg__SVGCSSRule__COLOR_PROFILE_RULE">COLOR_PROFILE_RULE</a> = 7;
};</pre><dl class="interface"><dt class="constants-header">Constants:</dt><dd><dl class="constants">
<dt id="__svg__SVGCSSRule__COLOR_PROFILE_RULE" class="constant first-child"><b>COLOR_PROFILE_RULE</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The rule is an <a href="http://www.w3.org/TR/SVG/color.html#InterfaceSVGColorProfileRule">@color-profile</a>.
</div></dd></dl></dd></dl>

<h3 id="InterfaceSVGRenderingIntent">4.5.33 Interface SVGRenderingIntent</h3>


 The <a class="idlinterface" href="types.html#InterfaceSVGRenderingIntent">SVGRenderingIntent</a> interface defines the enumerated list of
 possible values for <a href="color.html#ColorProfileElementRenderingIntentAttribute"><span class="attr-name">‘rendering-intent’</span></a> attributes or descriptors.
<pre class="idl">interface <b>SVGRenderingIntent</b> {
  // Rendering Intent Types
  const unsigned short <a href="types.html#__svg__SVGRenderingIntent__RENDERING_INTENT_UNKNOWN">RENDERING_INTENT_UNKNOWN</a> = 0;
  const unsigned short <a href="types.html#__svg__SVGRenderingIntent__RENDERING_INTENT_AUTO">RENDERING_INTENT_AUTO</a> = 1;
  const unsigned short <a href="types.html#__svg__SVGRenderingIntent__RENDERING_INTENT_PERCEPTUAL">RENDERING_INTENT_PERCEPTUAL</a> = 2;
  const unsigned short <a href="types.html#__svg__SVGRenderingIntent__RENDERING_INTENT_RELATIVE_COLORIMETRIC">RENDERING_INTENT_RELATIVE_COLORIMETRIC</a> = 3;
  const unsigned short <a href="types.html#__svg__SVGRenderingIntent__RENDERING_INTENT_SATURATION">RENDERING_INTENT_SATURATION</a> = 4;
  const unsigned short <a href="types.html#__svg__SVGRenderingIntent__RENDERING_INTENT_ABSOLUTE_COLORIMETRIC">RENDERING_INTENT_ABSOLUTE_COLORIMETRIC</a> = 5;
};</pre><dl class="interface"><dt class="constants-header">Constants in group “Rendering Intent Types”:</dt><dd><dl class="constants">
<dt id="__svg__SVGRenderingIntent__RENDERING_INTENT_UNKNOWN" class="constant first-child"><b>RENDERING_INTENT_UNKNOWN</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 The type is not one of predefined types. It is invalid to attempt to
 define a new value of this type or to attempt to switch an existing value
 to this type.

</div></dd>
<dt id="__svg__SVGRenderingIntent__RENDERING_INTENT_AUTO" class="constant"><b>RENDERING_INTENT_AUTO</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to a value of <span class="attr-value">'auto'</span>.

</div></dd>
<dt id="__svg__SVGRenderingIntent__RENDERING_INTENT_PERCEPTUAL" class="constant"><b>RENDERING_INTENT_PERCEPTUAL</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to a value of <span class="attr-value">'perceptual'</span>.

</div></dd>
<dt id="__svg__SVGRenderingIntent__RENDERING_INTENT_RELATIVE_COLORIMETRIC" class="constant"><b>RENDERING_INTENT_RELATIVE_COLORIMETRIC</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to a value of <span class="attr-value">'relative-colorimetric'</span>.

</div></dd>
<dt id="__svg__SVGRenderingIntent__RENDERING_INTENT_SATURATION" class="constant"><b>RENDERING_INTENT_SATURATION</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to a value of <span class="attr-value">'saturation'</span>.

</div></dd>
<dt id="__svg__SVGRenderingIntent__RENDERING_INTENT_ABSOLUTE_COLORIMETRIC" class="constant"><b>RENDERING_INTENT_ABSOLUTE_COLORIMETRIC</b><span class="idl-type-parenthetical"> (unsigned short)</span></dt><dd class="constant"><div>
 Corresponds to a value of <span class="attr-value">'absolute-colorimetric'</span>.

</div></dd></dl></dd></dl>

<div class="header bottom"><span class="namedate">SVG 1.1 (Second Edition) – 16 August 2011</span><a href="Overview.html">Top</a> ⋅ <a href="expanded-toc.html">Contents</a> ⋅ <a href="render.html">Previous</a> ⋅ <a href="struct.html">Next</a> ⋅ <a href="eltindex.html">Elements</a> ⋅ <a href="attindex.html">Attributes</a> ⋅ <a href="propidx.html">Properties</a></div><script src="style/expanders.js" type="text/javascript"> </script></body></html>
`
