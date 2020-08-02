package kubeconfig

import (
	"testing"

	"github.com/ahmetb/kubectx/internal/testutil"
)

func TestKubeconfig_GetCurrentContext(t *testing.T) {
	tl := WithMockKubeconfigLoader(`current-context: foo`)
	kc := new(Kubeconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	v := kc.GetCurrentContext()

	expected := "foo"
	if v != expected {
		t.Fatalf("expected=%q; got=%q", expected, v)
	}
}

func TestKubeconfig_GetCurrentContext_missingField(t *testing.T) {
	tl := WithMockKubeconfigLoader(`abc: def`)
	kc := new(Kubeconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	v := kc.GetCurrentContext()

	expected := ""
	if v != expected {
		t.Fatalf("expected=%q; got=%q", expected, v)
	}
}

func TestKubeconfig_UnsetCurrentContext(t *testing.T) {
	tl := WithMockKubeconfigLoader(testutil.KC().WithCurrentCtx("foo").ToYAML(t))
	kc := new(Kubeconfig).WithLoader(tl)
	if err := kc.Parse(); err != nil {
		t.Fatal(err)
	}
	if err := kc.UnsetCurrentContext(); err != nil {
		t.Fatal(err)
	}
	if err := kc.Save(); err != nil {
		t.Fatal(err)
	}

	out := tl.Output()
	expected := testutil.KC().WithCurrentCtx("").ToYAML(t)
	if out != expected {
		t.Fatalf("expected=%q; got=%q", expected, out)
	}
}
