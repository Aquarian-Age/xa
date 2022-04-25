#
# spec file for package chineseLunar
#
# Copyright (c) 2022 SUSE LLC
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via https://bugs.opensuse.org/
#


Name:           chineseLunar
Version:        1.0.9
Release:        0
Summary:        Constellation chineseLunar
License:        MIT
URL:            https://github.com/Aquarian-Age/xa
Source0:        chineseLunar.tar.gz

%if 0%{?suse_version} >= 1150
BuildRequires: golang(API) = 1.15
%else
BuildRequires: golang(API) = 1.17
%endif

BuildRequires:  gcc 
BuildRequires:  pkg-config 
BuildRoot:      %{_tmppath}/%{name}-%{version}-build   

%description
一个简单的中国农历
包含阳历 干支 阴历 二十八星宿 宗教节日鼠标悬停显示

%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/chineseLunar.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}

go build -o chineseLunar -mod=vendor -tags tempdll -ldflags="-s -w -X 'main.Version=1.0.9t' -X 'main.Mail=bGlhbmd6aTIwMjFAeWFuZGV4LmNvbQo='" -trimpath .

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -m0755 %{_builddir}/%{name}-%{version}/chineseLunar %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/chineseLunar

mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/chineseLunar/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}
install -m0755 %{_builddir}/%{name}-%{version}/chineseLunar.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/chineseLunar/chineseLunar.png
install -m0755 %{_builddir}/%{name}-%{version}/chineseLunar.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/chineseLunar.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/*.go
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/go.*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/chineseLunar*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/chineseLunar.png
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
%{_bindir}/chineseLunar
%{_datadir}/applications/chineseLunar.desktop
%{_datadir}/icons/chineseLunar
%{_datadir}/icons/chineseLunar/chineseLunar.png


%changelog
