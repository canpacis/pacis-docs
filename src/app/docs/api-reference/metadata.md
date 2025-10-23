# Metadata Package Documentation

## Overview

The `metadata` package provides a comprehensive set of Go structures and utilities for generating HTML metadata elements. It simplifies the creation of SEO tags, Open Graph protocol tags for social media sharing, Twitter Cards, and other standard HTML metadata.

This package is designed to work with the `github.com/canpacis/pacis/html` library and provides a type-safe way to generate all the metadata elements that should appear in an HTML document's `<head>` section.

## Installation

{noaccessory=true}
```go
import "github.com/canpacis/pacis/metadata"
```

## Core Types

### Metadata

The main structure that contains all metadata for an HTML document.

{noaccessory=true}
```go
type Metadata struct {
    Base            *url.URL
    Title           string
    Description     string
    Generator       string
    ApplicationName string
    Referrer        string
    Keywords        []string
    Authors         []Author
    Creator         string
    Publisher       string
    Canonical       string
    Alternates      []Alternate
    Robots          *Robots
    OpenGraph       *OpenGraph
    Twitter         *Twitter
}
```

**Fields:**

- `Base`: Base URL for all relative URLs in the document
- `Title`: Document title (appears in browser tabs and search results)
- `Description`: Brief description for SEO purposes
- `Generator`: Software that generated the page
- `ApplicationName`: Name of the web application
- `Referrer`: Controls Referer header behavior
- `Keywords`: List of relevant keywords
- `Authors`: List of document authors
- `Creator`: Content creator identifier
- `Publisher`: Content publisher identifier
- `Canonical`: Canonical URL to prevent duplicate content issues
- `Alternates`: Alternative versions of the page
- `Robots`: Search engine crawler directives
- `OpenGraph`: Open Graph metadata for social sharing
- `Twitter`: Twitter Card metadata

**Method:**

- `Node() html.Node`: Generates all HTML metadata elements as a fragment

### Author

Represents a document author with optional profile URL.

{noaccessory=true}
```go
type Author struct {
    Name string
    URL  string
}
```

**Fields:**

- `Name`: Author's full name
- `URL`: Optional link to author's profile or website

### Alternate

Represents alternative versions of the document (translations, mobile versions, etc.).

{noaccessory=true}
```go
type Alternate struct {
    Title    string
    Href     string
    HrefLang string
    Type     string
    Media    string
}
```

**Fields:**

- `Title`: Human-readable title of the alternate resource
- `Href`: URL of the alternate resource
- `HrefLang`: Language code (e.g., "en", "es", "fr")
- `Type`: MIME type of the alternate resource
- `Media`: Media query for the alternate resource

## Search Engine Optimization

### Robots

Controls search engine crawler behavior.

{noaccessory=true}
```go
type Robots struct {
    Index   bool
    Follow  bool
    Nocache bool
    Other   map[string]map[string]any
}
```

**Fields:**

- `Index`: Whether search engines should index the page
- `Follow`: Whether search engines should follow links on the page
- `Nocache`: Whether search engines should not cache the page
- `Other`: Additional directives for specific crawlers (e.g., "googlebot")

**Method:**

- `Node() html.Node`: Generates robots meta tags

**Example:**

{noaccessory=true}
```go
robots := &metadata.Robots{
    Index:   true,
    Follow:  true,
    Nocache: false,
    Other: map[string]map[string]any{
        "googlebot": {
            "max-image-preview": "large",
            "max-snippet":       -1,
        },
    },
}
```

## Social Media Integration

### OpenGraph

Open Graph protocol metadata for rich social media sharing (Facebook, LinkedIn, etc.).

{noaccessory=true}
```go
type OpenGraph struct {
    Title       string
    Description string
    URL         string
    SiteName    string
    Locale      string
    Type        string
    Images      []OpenGraphMedia
    Videos      []OpenGraphMedia
    Audio       []OpenGraphMedia
}
```

**Fields:**

- `Title`: Title for social shares
- `Description`: Brief description of the content
- `URL`: Canonical URL of the page
- `SiteName`: Name of the overall website
- `Locale`: Content locale (e.g., "en_US", "es_ES")
- `Type`: Content type (e.g., "website", "article", "video.movie")
- `Images`: Associated images
- `Videos`: Associated videos
- `Audio`: Associated audio files

**Method:**

- `Node() html.Node`: Generates Open Graph meta tags

**Example:**

{noaccessory=true}
```go
og := &metadata.OpenGraph{
    Title:       "My Amazing Article",
    Description: "Learn about amazing things",
    URL:         "https://example.com/article",
    SiteName:    "Example Site",
    Type:        "article",
    Images: []metadata.OpenGraphMedia{
        {
            URL:    "https://example.com/image.jpg",
            Width:  1200,
            Height: 630,
            Alt:    "Article thumbnail",
        },
    },
}
```

### OpenGraphMedia

Media content for Open Graph (images, videos, audio).

{noaccessory=true}
```go
type OpenGraphMedia struct {
    URL    string
    Width  int
    Height int
    Alt    string
}
```

**Fields:**

- `URL`: Location of the media file
- `Width`: Width in pixels (optional)
- `Height`: Height in pixels (optional)
- `Alt`: Alternative text (for images)

### Twitter

Twitter Card metadata for enhanced Twitter sharing.

{noaccessory=true}
```go
type Twitter struct {
    Card        string
    Title       string
    Description string
    SiteID      string
    Creator     string
    CreatorID   string
    Images      []string
}
```

**Fields:**

- `Card`: Card type ("summary", "summary_large_image", "player", "app")
- `Title`: Title for the Twitter card
- `Description`: Brief description
- `SiteID`: Twitter user ID of the website
- `Creator`: Twitter username of the creator (e.g., "@username")
- `CreatorID`: Twitter user ID of the creator
- `Images`: Image URLs for the card

**Method:**

- `Node() html.Node`: Generates Twitter Card meta tags

**Example:**

{noaccessory=true}
```go
twitter := &metadata.Twitter{
    Card:        "summary_large_image",
    Title:       "My Amazing Article",
    Description: "Learn about amazing things",
    Creator:     "@myhandle",
    Images:      []string{"https://example.com/image.jpg"},
}
```

## Usage Example

Here's a complete example of using the metadata package:

{noaccessory=true}
```go
package main

import (
    "net/url"
    "github.com/canpacis/pacis/metadata"
)

func main() {
    baseURL, _ := url.Parse("https://example.com")
    
    meta := &metadata.Metadata{
        Base:            baseURL,
        Title:           "My Awesome Website",
        Description:     "A comprehensive guide to awesome things",
        ApplicationName: "Awesome App",
        Keywords:        []string{"awesome", "guide", "tutorial"},
        Authors: []metadata.Author{
            {Name: "John Doe", URL: "https://example.com/john"},
        },
        Canonical: "https://example.com/page",
        
        Robots: &metadata.Robots{
            Index:  true,
            Follow: true,
        },
        
        OpenGraph: &metadata.OpenGraph{
            Title:       "My Awesome Website",
            Description: "A comprehensive guide to awesome things",
            URL:         "https://example.com/page",
            SiteName:    "Awesome Site",
            Type:        "website",
            Images: []metadata.OpenGraphMedia{
                {
                    URL:    "https://example.com/og-image.jpg",
                    Width:  1200,
                    Height: 630,
                    Alt:    "Site preview",
                },
            },
        },
        
        Twitter: &metadata.Twitter{
            Card:        "summary_large_image",
            Title:       "My Awesome Website",
            Description: "A comprehensive guide to awesome things",
            Images:      []string{"https://example.com/twitter-image.jpg"},
        },
    }
    
    // Generate HTML nodes
    metadataNode := meta.Node()
    
    // Use metadataNode in your HTML head section
}
```

## Generated HTML Output

The `Node()` methods generate appropriate HTML elements:

- **Title**: `<title>` element
- **Meta tags**: `<meta name="..." content="...">` for basic metadata
- **Link tags**: `<link rel="..." href="...">` for canonical URLs, alternates, and author links
- **Open Graph**: `<meta property="og:..." content="...">`
- **Twitter Cards**: `<meta name="twitter:..." content="...">`
- **Robots**: `<meta name="robots" content="...">`

## Best Practices

1. **Always set Title and Description**: These are crucial for SEO
2. **Use Canonical URLs**: Prevent duplicate content issues
3. **Provide Open Graph data**: Improves social media sharing appearance
4. **Set appropriate Robots directives**: Control crawler behavior
5. **Include image dimensions**: Improves social media preview loading
6. **Use descriptive Alt text**: Improves accessibility
7. **Set Twitter Card type appropriately**: "summary_large_image" for articles, "summary" for profiles
8. **Keep descriptions concise**: 150-160 characters for meta descriptions

## References

- [Open Graph Protocol](https://ogp.me/)
- [Twitter Cards Documentation](https://developer.twitter.com/en/docs/twitter-for-websites/cards)
- [HTML Meta Tags Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta)
- [Robots Meta Tag Specification](https://developers.google.com/search/docs/crawling-indexing/robots-meta-tag)