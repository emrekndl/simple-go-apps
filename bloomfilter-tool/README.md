# Bloom Filter Implementation in Go

Bu proje, Go programlama dili kullanılarak implemente edilmiş verimli bir **Bloom Filter** yapısını içerir. Bloom Filter, büyük veri kümelerinde bir elemanın varlığını test etmek için kullanılan, bellek açısından verimli bir olasılıksal veri yapısıdır.

## Bloom Filter Nedir?

Bloom Filter:
- **Olasılıksal bir veri yapısıdır** - yanlış pozitifler mümkündür ama yanlış negatifler imkânsızdır
- **Bellek verimlidir** - hash tablolarına kıyasla çok daha az bellek kullanır
- **Temel kullanım alanları**:
  - Önbellek filtreleme
  - Yazım denetleyicileri
  - Güvenlik uygulamaları
  - Büyük veritabanı sorgu optimizasyonu

## Teknik Parametreler ve Formüller

Bloom Filter performansı 4 temel parametreye bağlıdır:

- **n**: Filtreye eklenen eleman sayısı (örn: 5683 kelime)
- **p**: Yanlış pozitif olasılığı (örn: %0.1 = 0.001)
- **m**: Bit dizisinin uzunluğu (bit cinsinden)
- **k**: Kullanılan hash fonksiyonu sayısı

- Bloom Filter hesaplamaları: [bloom-filter-calc](https://hur.st/bloomfilter/)

Optimal parametre hesaplamaları:

```go
// Optimal bit dizisi boyutu (m)
m = ceil((n * log(p)) / log(1 / pow(2, log(2))))

// Optimal hash fonksiyonu sayısı (k)
k = round((m / n) * log(2))

// Yanlış pozitif olasılığı (p)
p = pow(1 - exp(-k / (m / n)), k)
```

Örnek değerlerimiz:
- n = 5683 (kelime)
- p = 0.001 (%0.1)
- m = 81708 bit (~10.2KB)
- k = 10 hash fonksiyonu

## Kurulum ve Çalıştırma

1. **Gereksinimler**: Sisteminizde [Go](https://golang.org/dl/) yüklü olmalı (v1.16+)

2. **Kodu indirin**:
   ```bash
   git clone https://github.com/simple-go-apps/bloom-filter-tool.git
   cd bloom-filter-tool
   ```

3. **Programı çalıştırın**:
   ```bash
   go run main.go
   ```

4. **Bloom Filter'ı JSON olarak dışa aktarma** (opsiyonel):
   ```go
   // main.go içindeki bu satırın yorumunu kaldırın:
   bf.ExportToJSON("bloom_data.json", k)
   ```
5. **Bloom Filter'ı JSON'dan Yükleme (opsiyonel)
    ```go
    bf, err := ImportFromJSON("bloom_data.json")
    if err != nil {
        log.Println("Hata:", err)
        // Yeni filtre oluştur
    }
```
## Örnek Çıktı

```
Adding item: bloom
  Hash 0 -> index: 1234 (byte: 154, bit: 2)
  Hash 1 -> index: 5678 (byte: 709, bit: 6)
  ...
Adding item: filter
  ...

Checking item: dünya
  Hash 0 -> index: 2345 (byte: 293, bit: 1) = true
  Hash 1 -> index: 6789 (byte: 848, bit: 5) = true
  ...
Result: true

Checking item: xqdmw
  Hash 0 -> index: 3456 (byte: 432, bit: 0) = false
  Result: false
```

## Performans Analizi

| Parametre | Değer | Açıklama |
|-----------|-------|----------|
| n | 5,683 | Kelime listesindeki öğe sayısı |
| m | 81,708 bit | Bit dizisinin boyutu (~10.2KB) |
| k | 10 | Kullanılan hash fonksiyonu sayısı |
| Teorik p | 0.001 | Beklenen yanlış pozitif oranı |
| Gerçek p | ~0.0009 | Pratik testlerde gözlemlenen oran |

## Kullanım Senaryoları

Bu implementasyon ideal olarak şu durumlarda kullanılabilir:
- Büyük kelime listeleri için yazım denetleyici
- Parola denetimi sistemleri
- Önbelleğe alınmamış sorguların tespiti
- URL filtreleme sistemleri

## Lisans

Bu proje [MIT Lisansı](LICENSE) altında lisanslanmıştır.
