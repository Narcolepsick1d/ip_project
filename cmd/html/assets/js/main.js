gsap.from(".home__img", { opacity: 0, duration: 1, delay: 0.8, x: 60 });
gsap.from(".home__data", { opacity: 0, duration: 1, delay: 0.8, y: 25 });
gsap.from(".home__welcome, .home__name, .home__description, .home__button", {
  opacity: 0,
  duration: 2,
  delay: 1,
  y: 25,
  ease: "expo.out",
  stagger: 1,
});

gsap.from(".nav__logo, .nav__toggle", {
  opacity: 0,
  duration: 2,
  delay: 1,
  y: 25,
  ease: "expo.out",
  stagger: 1,
});
gsap.from(".nav__item", {
  opacity: 0,
  duration: 2,
  delay: 1,
  y: 25,
  ease: "expo.out",
  stagger: 1,
});
gsap.from(".home__social-icon", {
  opacity: 0,
  duration: 2,
  delay: 2.3,
  y: 25,
  ease: "expo.out",
  stagger: 1,
});

gsap.from(
  ".about__img, .about__data, .about__description,.section-subtitle,.section-title,.about__number, .about__achievement",
  { opacity: 0, duration: 2, delay: 2.5, y: 25, ease: "expo.out", stagger: 0.2 }
);

gsap.from(".properties__item, .properties__data, .properties__content", {
  opacity: 0,
  duration: 2,
  delay: 2.5,
  y: 25,
  ease: "expo.out",
  stagger: 1,
});
gsap.from(".contact__box, .contact__title, .contact__description", {
  opacity: 0,
  duration: 2,
  delay: 2.5,
  y: 25,
  ease: "expo.out",
  stagger: ``,
});

const showMenu = (toggleId, navId) => {
  const toggle = document.getElementById(toggleId),
    nav = document.getElementById(navId);

  if (toggle && nav) {
    toggle.addEventListener("click", () => {
      nav.classList.toggle("show-menu");
    });
  }
};
showMenu("nav-toggle", "nav-menu");




/*===== SCROLL SECTIONS ACTIVE LINK =====*/
const sections = document.querySelectorAll("section[id]");

function scrollActive() {
  const scrollY = window.pageYOffset;

  sections.forEach((current) => {
    const sectionHeight = current.offsetHeight;
    const sectionTop = current.offsetTop - 50;
    sectionId = current.getAttribute("id");

    if (scrollY > sectionTop && scrollY <= sectionTop + sectionHeight) {
      document
        .querySelector(".nav__menu a[href*=" + sectionId + "]")
        .classList.add("active-link");
    } else {
      document
        .querySelector(".nav__menu a[href*=" + sectionId + "]")
        .classList.remove("active-link");
    }
  });
}
window.addEventListener("scroll", scrollActive);

/*===== CHANGE BACKGROUND HEADER =====*/
function scrollHeader() {
  const header = document.getElementById("header");
  if (this.scrollY >= 200) header.classList.add("scroll-header");
  else header.classList.remove("scroll-header");
}
window.addEventListener("scroll", scrollHeader);

/*===== MIXITUP FILTER PORTFOLIO =====*/
const mixer = mixitup(".properties__container", {
  selectors: {
    target: ".properties__content",
  },
  animation: {
    duration: 400,
  },
});
