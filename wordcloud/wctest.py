from wordcloud import WordCloud
from matplotlib import pyplot as plt
import time
import MeCab

with open("inja.txt", mode="rt", encoding="utf-8") as f:
  cloud_text = f.read()

tagger = MeCab.Tagger("")
tagger.parse("")

start = time.time()
node = tagger.parseToNode(cloud_text)

wl = []
while node:
  word_type = node.feature.split(",")[0]
  if word_type == "名詞":
    wl.append(node.surface)
  node = node.next

word_chain = " ".join(wl)

#print (word_chain)
word_cloud = WordCloud(
  font_path="./font/SourceHanSerifK-Light.otf",
  background_color="white",
  colormap="tab20"
).generate(word_chain)
elapsed = time.time() - start

fig = plt.figure()
plt.imshow(word_cloud)
plt.axis("off")
fig.savefig("img.png")
print("elapsed : " + str(elapsed))

#plt.axis("off")
#plt.show()
#