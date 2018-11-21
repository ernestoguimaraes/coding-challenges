import ascii.createAdidasAsciiLogo
import org.junit.Test
import kotlin.test.assertEquals


class TestSource {
    @Test
    fun testWidth2() {
        val expected = """
            @@
          @@ @@
        @@ @@ @@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }

    @Test
    fun testWidth3() {
        val expected = """
              @@@
               @@@
           @@@  @@@
            @@@  @@@
        @@@  @@@  @@@
         @@@  @@@  @@@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }

    @Test
    fun testWidth5() {
        val expected = """
                  @@@@@
                   @@@@@
             @@@@@  @@@@@
              @@@@@  @@@@@
        @@@@@  @@@@@  @@@@@
         @@@@@  @@@@@  @@@@@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }

        @Test
    fun testWidth7() {
        val expected = """
                      @@@@@@@
                       @@@@@@@
                        @@@@@@@
               @@@@@@@   @@@@@@@
                @@@@@@@   @@@@@@@
                 @@@@@@@   @@@@@@@
        @@@@@@@   @@@@@@@   @@@@@@@
         @@@@@@@   @@@@@@@   @@@@@@@
          @@@@@@@   @@@@@@@   @@@@@@@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }

            @Test
    fun testWidth9() {
        val expected = """
                          @@@@@@@@@
                           @@@@@@@@@
                            @@@@@@@@@
                 @@@@@@@@@   @@@@@@@@@
                  @@@@@@@@@   @@@@@@@@@
                   @@@@@@@@@   @@@@@@@@@
        @@@@@@@@@   @@@@@@@@@   @@@@@@@@@
         @@@@@@@@@   @@@@@@@@@   @@@@@@@@@
          @@@@@@@@@   @@@@@@@@@   @@@@@@@@@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }

                @Test
    fun testWidth16() {
        val expected = """
                                         @@@@@@@@@@@@@@@@
                                         @@@@@@@@@@@@@@@@
                                          @@@@@@@@@@@@@@@@
                                           @@@@@@@@@@@@@@@@
                        @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
                         @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
                          @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
                           @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
        @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
         @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
          @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
           @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@    @@@@@@@@@@@@@@@@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }

                    @Test
    fun testWidth21() {
        val expected = """
                                                  @@@@@@@@@@@@@@@@@@@@@
                                                   @@@@@@@@@@@@@@@@@@@@@
                                                    @@@@@@@@@@@@@@@@@@@@@
                                                     @@@@@@@@@@@@@@@@@@@@@
                                                      @@@@@@@@@@@@@@@@@@@@@
                             @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
                              @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
                               @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
                                @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
                                 @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
        @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
         @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
          @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
           @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
            @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@
        """.trim()
        assertEquals(expected, createAdidasAsciiLogo(2))
    }
}
