---
#  "Exhibition Autocomplete: Displaying imagesets from The Curator"
### 2020-09-05 {#date}
### word2vec, imagesets, webapps {#tags}
---

In the [last](https://lancewakeling.net/blog/2020-07-04-the-curator/) [three](https://lancewakeling.net/blog/2020-07-16-the-curator-2/) [posts](https://lancewakeling.net/blog/2020-08-02-the-curator-3/) I focused on the data side of a personal project, The Curator. For this post, I'll quickly document the web interface that I'm calling [Exhibition Autocomplete](https://github.com/disposableraft/the-curator-webapp).

The basic user experience is: 

1. Begin typing and select a term from the autocomplete list;
2. Browse a collection of images;
3. Rinse and repeat.

![screenshot from github](https://raw.githubusercontent.com/disposableraft/the-curator-webapp/6d95c5c4f70a0bc272b1c7d17469d005c4157ef8/public/curator-v0.1.gif)

Maybe it's because I built it, but the interface is a pleasure to play around with. The collections are composed of beautiful images that *seem* to fit together. 

The model deployed, [`1.8.2`](https://github.com/disposableraft/the-curator#benchmarks), has a 73% error rate for predicting a similar artist within the same art historical category. Though the error rate is higher than 50%, the image collections do not look random to my eye. That is, Impressionists are not *too* mixed up with Minimalists and Conceptualists. 

![Artists similar to Van Gogh](/images/exhib-auto-van-gogh.png)

<Caption>Artworks by artists similar to Vincent Van Gogh. Mondrian — row two, column two — is grouped with artists similar to Van Gogh.</Caption>

![Artists similar to Helen Levitt](/images/exhib-auto-helen-levitt.png)

<Caption>Artworks by artists similar to Helen Levitt. The imageset is almost entirely comprised of black-and-white photographs.</Caption>

The dataflow works something like the following figure illustrates. 

![Data flow](/images/exhib-auto-data-flow.png)

The python model is pretty slow. But all the outputs from this call can be memoized, so convenience is the only reason to use the python model instead of a lookup table in this case.

```bash
> time python python/resolver.py "Pablo Picasso"
{"name": "Pablo Picasso", "artists": ["Henri Matisse", "Amedeo Modigliani", "Alberto Giacometti", "Raoul Dufy", "Franti\u009aek Kupka", "Lovis Corinth", "Alexander Calder", "Henry Moore", "Odilon Redon", "Ernst Ludwig Kirchner", "Jules Pascin", "Auguste Rodin"]}

python python/resolver.py "Pablo Picasso"  0.79s user 0.21s system 69% cpu 1.461 total
```
Anyway, not sure yet if I'll publish the site online. The images are gathered from a Google custom search engine, which restricts free accounts to a low quota. I got around this by writing a really simple cache.

```javascript
// Attempt to get a cached result.
getCachedResult(name, (err, cache) => {
  // If one isn't found submit a search request.
  if (err) {
    search(name, serverRuntimeConfig)
      .then((data) => {
        if (data.error) {
          res.json(data);
        } else {
          // Cache the search results.
          setCachedResult(name, data);
          res.json(data.items[0]);
        }
      })
      .catch((err) => {
        res.json({ error: err });
      });
  } else {
    // Otherwise, return the existing cache data.
    res.json(JSON.parse(cache).items[0]);
  }
});
```

So, with a lookup table, loading animations and a more sophisticated cache, it would be possible to go live...
